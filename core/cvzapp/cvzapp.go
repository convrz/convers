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

package cvzapp

import "go.uber.org/fx"

var options = fx.Provide()

func Provide(constructors ...interface{}) {
	options = fx.Options(options, fx.Provide(constructors...))
}

func Option() fx.Option {
	return options
}

func Invoke(constructors ...interface{}) {
	options = fx.Options(options, fx.Invoke(constructors...))
}

type App interface {
	Run() error
}