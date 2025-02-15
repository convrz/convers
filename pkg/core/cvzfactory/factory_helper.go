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

package cvzfactory

import (
	"context"

	"github.com/convrz/convers/pkg/core/cvzapp"
	cvzinternal "github.com/convrz/convers/pkg/core/internal"

	"github.com/convrz/convers/pkg/logger"

	"go.uber.org/fx"
)

func _start(srv cvzapp.Server) {
	logger.Fatal(srv.ListenAndServe())
}

func _main(lc fx.Lifecycle, srv cvzapp.Server) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go _start(srv)
			return nil
		},
		OnStop: func(_ context.Context) error {
			return nil
		},
	})
}

// Build builds the application.
func Build(constructor interface{}) cvzapp.Application {
	cvzinternal.Provide(constructor)

	return &container{
		engine: fx.New(cvzinternal.Option(), fx.Invoke(_main)),
	}
}
