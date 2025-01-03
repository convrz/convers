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

package cvzhook

import (
	"context"
	"github.com/convrz/convers/core/cvzapp"
	"go.uber.org/fx"
)

func UseWithLifeCycle[T any](constructor func() T, onStart func(T) error, onStop func(T) error) func() T {
	decorateConstructor := func(lc fx.Lifecycle) T {
		ins := constructor()

		lc.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return onStart(ins)
			},
			OnStop: func(ctx context.Context) error {
				return onStop(ins)
			},
		})

		return ins
	}

	cvzapp.Provide(decorateConstructor)

	return constructor
}

func Use[fn any](constructor fn) fn {
	cvzapp.Provide(constructor)
	return constructor
}

func UseBefore(fn func(ctx context.Context) error) {
	function := func(lc fx.Lifecycle) {
		lc.Append(fx.Hook{
			OnStart: fn,
		})
	}

	cvzapp.Invoke(function)
}

func UseAfter(fn func(ctx context.Context) error) {
	function := func(lc fx.Lifecycle) {
		lc.Append(fx.Hook{
			OnStop: fn,
		})
	}

	cvzapp.Invoke(function)
}