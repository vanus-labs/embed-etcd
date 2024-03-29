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
	etcd := embedetcd.New(map[string]string{})
	ctx := context.Background()
	err := etcd.Init(ctx, embedetcd.Config{
		Name:                "etcd-1",
		DataDir:             "/Users/wenfeng/tmp/embed_etcd/node1",
		ListenClientAddr:    "0.0.0.0:2379",
		ListenPeerAddr:      "0.0.0.0:2380",
		AdvertiseClientAddr: "localhost:2279,127.0.0.1:2279",
		AdvertisePeerAddr:   "localhost:2380,127.0.0.1:2380",
		Clusters:            []string{"etcd-1=http://localhost:2380,etcd-2=http://localhost:3380,etcd-3=http://localhost:4380"},
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
	etcd := embedetcd.New(map[string]string{})
	ctx := context.Background()
	err := etcd.Init(ctx, embedetcd.Config{
		Name:                "etcd-2",
		DataDir:             "/Users/wenfeng/tmp/embed_etcd/node2",
		ListenClientAddr:    "0.0.0.0:3379",
		ListenPeerAddr:      "0.0.0.0:3380",
		AdvertiseClientAddr: "localhost:3379,127.0.0.1:3379",
		AdvertisePeerAddr:   "localhost:3380,127.0.0.1:3380",
		Clusters:            []string{"etcd-1=http://localhost:2380,etcd-2=http://localhost:3380,etcd-3=http://localhost:4380"},
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
	etcd := embedetcd.New(map[string]string{})
	ctx := context.Background()
	err := etcd.Init(ctx, embedetcd.Config{
		Name:                "etcd-3",
		DataDir:             "/Users/wenfeng/tmp/embed_etcd/node3",
		ListenClientAddr:    "0.0.0.0:4379",
		ListenPeerAddr:      "0.0.0.0:4380",
		AdvertiseClientAddr: "localhost:4379,127.0.0.1:4379",
		AdvertisePeerAddr:   "localhost:4380,127.0.0.1:4380",
		Clusters:            []string{"etcd-1=http://localhost:2380,etcd-2=http://localhost:3380,etcd-3=http://localhost:4380"},
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
