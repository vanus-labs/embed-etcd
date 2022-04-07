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
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	etcd := New()
	ctx := context.Background()
	cfg := Config{
		Name:        "test-1",
		DataDir:     "/tmp/embed_etcd",
		ClientAddrs: "0.0.0.0:2379",
		PeerAddrs:   "0.0.0.0:2380",
	}
	if err := etcd.Init(ctx, cfg); err != nil {
		panic(err)
	}
	if _, err := etcd.Start(ctx); err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Hour)
}
