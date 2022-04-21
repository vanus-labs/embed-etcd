// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package embedetcd

import (
	"context"
	"errors"
	"github.com/linkall-labs/embed-etcd/log"
	"go.etcd.io/etcd/client/pkg/v3/types"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	etcdTimeout = time.Second * 3
	httpSchema  = "http://"
	httpsSchema = "https://"
)

var (
	ErrStartEtcd         = errors.New("start etcd failed")
	ErrStartEtcdCanceled = errors.New("etcd start canceled")
)

type EmbedEtcd interface {
	Init(context.Context, Config) error
	Start(context.Context) (<-chan struct{}, error)
	Stop(context.Context)
}

func New(topology map[string]string) *embedEtcd {
	ee := &embedEtcd{
		topology: topology,
	}
	ee.ctx, ee.cancel = context.WithCancel(context.Background())
	return ee
}

type MembershipEventProcessor func(ctx context.Context, event MembershipChangedEvent) error

type embedEtcd struct {
	instance   *embed.Etcd
	cfg        *Config
	etcdCfg    *embed.Config
	ctx        context.Context
	cancel     context.CancelFunc
	client     *clientv3.Client
	httpClient *http.Client
	isLeader   bool
	mutex      sync.RWMutex
	once       sync.Once
	handlers   []MembershipEventProcessor
	topology   map[string]string
}

func (ee *embedEtcd) Init(ctx context.Context, cfg Config) error {
	ee.cfg = &cfg
	ee.etcdCfg = embed.NewConfig()
	ee.etcdCfg.Logger = "zap"
	ee.etcdCfg.LogLevel = "warn"
	ee.etcdCfg.LogOutputs = []string{"stdout"}
	ee.etcdCfg.Name = cfg.Name
	ee.etcdCfg.Dir = cfg.DataDir

	var err error
	if ee.etcdCfg.LCUrls, err = types.NewURLs([]string{httpSchema + ee.cfg.ClientAddr}); err != nil {
		return err
	}

	if ee.etcdCfg.LPUrls, err = types.NewURLs([]string{httpSchema + ee.cfg.PeerAddr}); err != nil {
		return err
	}

	if ee.etcdCfg.ACUrls, err = types.NewURLs([]string{httpSchema + ee.cfg.ClientAddr}); err != nil {
		return err
	}

	if ee.etcdCfg.APUrls, err = types.NewURLs([]string{httpSchema + ee.cfg.PeerAddr}); err != nil {
		return err
	}

	ee.etcdCfg.InitialCluster = strings.Join(ee.cfg.Clusters, ",")
	return nil
}

func (ee *embedEtcd) Start(ctx context.Context) (<-chan struct{}, error) {
	e, err := embed.StartEtcd(ee.etcdCfg)
	if err != nil {
		return nil, err
	}
	ee.instance = e
	select {
	case <-e.Server.ReadyNotify():
		log.Info("etcd server started", map[string]interface{}{
			"name":        ee.cfg.Name,
			"client_addr": ee.cfg.ClientAddr,
			"peer_addr":   ee.cfg.PeerAddr,
		})
	}
	go ee.run()
	return e.Server.StoppingNotify(), nil
}

func (ee *embedEtcd) Stop(ctx context.Context) {
	ee.instance.Server.Stop()
}

func (ee *embedEtcd) run() {
	for {
		select {
		case <-ee.ctx.Done():
			log.Info("membership changed listener stopped", nil)
			return
		default:
		}
		event := MembershipChangedEvent{
			Leader: ee.instance.Server.Lead(),
		}

		isLeader := ee.instance.Server.Leader().String() == ee.instance.Server.ID().String()
		var err error
		roleChanged := false
		if ee.isLeader && !isLeader {
			ee.isLeader = false
			event.Type = EventBecomeFollower
			roleChanged = true
		} else if !ee.isLeader && isLeader {
			ee.isLeader = true
			event.Type = EventBecomeLeader
			roleChanged = true
		}
		if roleChanged {
			log.Info("the role of node changed", map[string]interface{}{
				"role": event.Type,
			})
			ee.mutex.RLock()
			for _, handler := range ee.handlers {
				err = handler(ee.ctx, event)
				if err != nil {
					break
				}
			}
			ee.mutex.RUnlock()
			if err != nil {
				ee.isLeader = false
				ee.ResignIfLeader()
				log.Error("failed to process membership event", map[string]interface{}{
					"error": err,
				})
				return
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func (ee *embedEtcd) RegisterMembershipChangedProcessor(handler MembershipEventProcessor) {
	ee.mutex.Lock()
	defer ee.mutex.Unlock()
	ee.handlers = append(ee.handlers, handler)
}

func (ee *embedEtcd) ResignIfLeader() {
	if err := ee.instance.Server.TransferLeadership(); err != nil {
		log.Warning("transfer leadership to another failed", map[string]interface{}{
			"error": err,
		})
	}
}

func (ee *embedEtcd) IsLeader() bool {
	return ee.isLeader && ee.instance.Server.Leader().String() == ee.instance.Server.ID().String()
}

func (ee *embedEtcd) GetLeaderID() string {
	return ee.instance.Server.Leader().String()
}

func (ee *embedEtcd) GetLeaderAddr() string {
	id := ee.instance.Server.Leader()
	if id == 0 {
		return "unavailable"
	}
	member := ee.instance.Server.Cluster().Member(id)
	if member == nil {
		return "unavailable"
	}
	if len(member.ClientURLs) == 0 {
		return "unavailable"
	}
	urls := member.PeerURLs
	for _, v := range ee.cfg.Clusters {
		strs := strings.Split(v, "=")
		if strs[1] == urls[0] {
			return ee.topology[strs[0]]
		}
	}
	return ""
}

var (
	// EtcdStartTimeout the timeout of the startup etcd.
	EtcdStartTimeout = time.Minute * 5
)

func (ee *embedEtcd) startEtcd(ctx context.Context) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, EtcdStartTimeout)
	defer cancel()

	etcd, err := embed.StartEtcd(ee.etcdCfg)
	if err != nil {
		return ErrStartEtcd
	}

	// TODO support add TLS
	select {
	// Wait etcd until it is ready to use
	case <-etcd.Server.ReadyNotify():
	case <-timeoutCtx.Done():
		return ErrStartEtcdCanceled
	}
	return nil
}
