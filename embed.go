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

package embed_etcd

import (
	"context"
	error2 "github.com/linkall-labs/embed-etcd/errors"
	"github.com/linkall-labs/embed-etcd/etcdutil"
	"github.com/linkall-labs/embed-etcd/log"
	"go.etcd.io/etcd/client/pkg/v3/types"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/server/v3/embed"
	"go.uber.org/zap"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	etcdTimeout = time.Second * 3
)

type EmbedEtcd interface {
	Init(context.Context, Config) error
	Start(context.Context) (<-chan error, error)
	Stop(context.Context) error
	AddMember(context.Context, ...string) error
	RemoveMember(context.Context, uint64) error
	IsLeader(context.Context) bool
}

func New() EmbedEtcd {
	return &embedEtcd{
		ctx:        context.Background(),
		exitSignal: make(chan error, 0),
	}
}

type embedEtcd struct {
	instance   *embed.Etcd
	cfg        *Config
	etcdCfg    *embed.Config
	ctx        context.Context
	exitSignal chan error
	client     *clientv3.Client
	httpClient *http.Client
}

func (ee *embedEtcd) Init(ctx context.Context, cfg Config) error {
	ee.etcdCfg = embed.NewConfig()
	ee.etcdCfg.Logger = "zap"
	ee.etcdCfg.LogOutputs = []string{"stdout"}
	ee.etcdCfg.Name = cfg.Name
	ee.etcdCfg.Dir = cfg.DataDir
	return nil
}

func (ee *embedEtcd) Start(ctx context.Context) (<-chan error, error) {
	e, err := embed.StartEtcd(ee.etcdCfg)
	if err != nil {
		return nil, err
	}
	ee.instance = e
	return ee.exitSignal, nil
}

func (ee *embedEtcd) Stop(ctx context.Context) error {
	ee.exitSignal <- nil
	return nil
}

func (ee *embedEtcd) AddMember(ctx context.Context, urls ...string) error {
	return nil
}

func (ee *embedEtcd) RemoveMember(ctx context.Context, id uint64) error {
	return nil
}

func (ee *embedEtcd) IsLeader(ctx context.Context) bool {
	return false
}

var (
	// EtcdStartTimeout the timeout of the startup etcd.
	EtcdStartTimeout = time.Minute * 5
)

func (ee *embedEtcd) startEtcd(ctx context.Context) error {
	newCtx, cancel := context.WithTimeout(ctx, EtcdStartTimeout)
	defer cancel()

	etcd, err := embed.StartEtcd(ee.etcdCfg)
	if err != nil {
		return error2.ErrStartEtcd
	}

	// Check cluster ID
	urlMap, err := types.NewURLsMap(ee.etcdCfg.InitialCluster)
	if err != nil {
		return error2.ErrEtcdURLMap
	}

	//tlsConfig, err := ee.cfg.Security.ToTLSConfig()
	//if err != nil {
	//	return err
	//}

	if err = etcdutil.CheckClusterID(etcd.Server.Cluster().ID(), urlMap, nil); err != nil {
		return err
	}

	select {
	// Wait etcd until it is ready to use
	case <-etcd.Server.ReadyNotify():
	case <-newCtx.Done():
		return error2.ErrStartEtcdCanceled
	}

	endpoints := []string{ee.etcdCfg.ACUrls[0].String()}
	log.Info("create etcd v3 client", map[string]interface{}{
		"endpoints": endpoints,
		//"cert":      ee.cfg.Security,
	})

	lgc := zap.NewProductionConfig()
	lgc.Encoding = log.ZapEncodingName
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: etcdTimeout,
		//TLS:         tlsConfig,
		LogConfig: &lgc,
	})
	if err != nil {
		return error2.ErrNewEtcdClient
	}

	etcdServerID := uint64(etcd.Server.ID())

	//update advertise peer urls.
	etcdMembers, err := etcdutil.ListEtcdMembers(client)
	if err != nil {
		return err
	}
	for _, m := range etcdMembers.Members {
		if etcdServerID == m.ID {
			etcdPeerURLs, err := parseUrls(m.PeerURLs...)
			if err != nil {
				return err
			}

			if isEqual(ee.etcdCfg.APUrls, etcdPeerURLs) {
				log.Info("update advertise peer urls", map[string]interface{}{
					"from": ee.etcdCfg.APUrls,
					"to":   etcdPeerURLs,
				})
				ee.etcdCfg.APUrls = etcdPeerURLs
			}
		}
	}
	ee.client = client
	ee.httpClient = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true,
			//TLSClientConfig:   tlsConfig,
		},
	}

	//s.member = member.NewMember(etcd, client, etcdServerID)
	return nil
}

func isEqual(urls1 []url.URL, urls2 []url.URL) bool {
	if len(urls1) != len(urls2) {
		return false
	}
	sort.Slice(urls1, func(i, j int) bool {
		return strings.Compare(urls1[i].String(), urls1[j].String()) > 0
	})
	sort.Slice(urls2, func(i, j int) bool {
		return strings.Compare(urls2[i].String(), urls2[j].String()) > 0
	})
	for idx := 0; idx < len(urls1); idx++ {
		if urls1[idx].String() != urls2[idx].String() {
			return false
		}
	}
	return true
}

// parseUrls parse a string into multiple urls.
func parseUrls(items ...string) ([]url.URL, error) {
	urls := make([]url.URL, 0, len(items))
	for _, item := range items {
		u, err := url.Parse(item)
		if err != nil {
			return nil, error2.ErrURLParse
		}

		urls = append(urls, *u)
	}

	return urls, nil
}
