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

	cvzinternal "github.com/convrz/convers/core/internal"
	"github.com/convrz/convers/pkg/logger"

	"github.com/convrz/convers/core/cvzapp"

	"go.uber.org/fx"
)

func _main(lc fx.Lifecycle, app cvzapp.Application) {
	start := func() {
		logger.Fatal(app.Run())
	}

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go start()
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
