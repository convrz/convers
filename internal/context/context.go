/*
Copyright 2024 The Convērs Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package context

import (
	"context"
	"time"

	"github.com/convrz/convers/api/types/v1"
	"github.com/convrz/convers/pkg/protobuf/proto"
)

const contextKey = 0

func New(ctxBase context.Context, msg *types.Context, deadline time.Time) (context.Context, context.CancelFunc) {
	if msg == nil {
		return context.WithDeadline(ctxBase, deadline)
	}

	msgBin, _ := proto.Marshal(msg)

	ctx := context.WithValue(ctxBase, contextKey, msgBin)

	return context.WithDeadline(ctx, deadline)
}

func ValueOf(ctx context.Context) (*types.Context, bool) {
	msgBin, ok := ctx.Value(contextKey).([]byte)
	if !ok {
		return nil, false
	}

	var msg types.Context
	if err := proto.Unmarshal(msgBin, &msg); err != nil {
		return nil, false
	}

	return &msg, true
}
