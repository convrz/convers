/*
 * Copyright 2024 The Convērs Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package cvzhook provides a way to hook into the application lifecycle.
package cvzhook

import (
	"context"

	cvzinternal "github.com/convrz/convers/core/internal"
	"go.uber.org/fx"
)

// UseBefore uses the given function before the application starts.
func UseBefore(fn func(ctx context.Context) error) {
	function := func(lc fx.Lifecycle) {
		lc.Append(fx.Hook{
			OnStart: fn,
		})
	}

	cvzinternal.Invoke(function)
}

// UseAfter adds a hook to be executed after the application has stopped.
func UseAfter(fn func(ctx context.Context) error) {
	function := func(lc fx.Lifecycle) {
		lc.Append(fx.Hook{
			OnStop: fn,
		})
	}

	cvzinternal.Invoke(function)
}
