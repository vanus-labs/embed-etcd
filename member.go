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
)

type EventType string

const (
	EventBecomeLeader   = "become_leader"
	EventBecomeFollower = "become_follower"
)

type MembershipChangedEvent struct {
	Type   EventType
	Leader uint64
}

type Member interface {
	RegisterMembershipChangedProcessor(context.Context, func(context.Context, MembershipChangedEvent) error)
	ResignIfLeader(context.Context)
	IsLeader() bool
	GetLeaderID() string
	GetLeaderAddr() string
}

type LeaderInfo struct {
	LeaderID   string `json:"leader_id"`
	LeaderName string `json:"leader_name"`
}
