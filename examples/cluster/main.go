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

package main

import (
	"context"
	embedetcd "github.com/linkall-labs/embed-etcd"
	"time"
)

func main() {
	go startA()
	go startB()
	go startC()
	time.Sleep(time.Hour)
}

func startA() {
	etcd := embedetcd.New()
	ctx := context.Background()
	err := etcd.Init(ctx, embedetcd.Config{
		Name:       "etcd-1",
		DataDir:    "/Users/wenfeng/tmp/embed_etcd/node1",
		ClientAddr: "localhost:2379",
		PeerAddr:   "localhost:2380",
		Clusters:   []string{"etcd-1=http://localhost:2380,etcd-2=http://localhost:3380,etcd-3=http://localhost:4380"},
	})
	if err != nil {
		panic(err)
	}
	_, err = etcd.Start(ctx)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Hour)
}

func startB() {
	etcd := embedetcd.New()
	ctx := context.Background()
	err := etcd.Init(ctx, embedetcd.Config{
		Name:       "etcd-2",
		DataDir:    "/Users/wenfeng/tmp/embed_etcd/node2",
		ClientAddr: "localhost:3379",
		PeerAddr:   "localhost:3380",
		Clusters:   []string{"etcd-1=http://localhost:2380,etcd-2=http://localhost:3380,etcd-3=http://localhost:4380"},
	})
	if err != nil {
		panic(err)
	}
	_, err = etcd.Start(ctx)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Hour)
}

func startC() {
	etcd := embedetcd.New()
	ctx := context.Background()
	err := etcd.Init(ctx, embedetcd.Config{
		Name:       "etcd-3",
		DataDir:    "/Users/wenfeng/tmp/embed_etcd/node3",
		ClientAddr: "localhost:4379",
		PeerAddr:   "localhost:4380",
		Clusters:   []string{"etcd-1=http://localhost:2380,etcd-2=http://localhost:3380,etcd-3=http://localhost:4380"},
	})
	if err != nil {
		panic(err)
	}
	_, err = etcd.Start(ctx)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Hour)
}
