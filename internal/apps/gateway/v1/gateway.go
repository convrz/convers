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

package gateway

import (
	"context"
	"github.com/convrz/convers/core/cvzapp"
	"github.com/convrz/convers/core/cvzruntime"
	_ "github.com/convrz/convers/internal/apps/gateway/v1/init"
	"github.com/convrz/convers/internal/apps/gateway/v1/registry"
	"log"
)

func New() cvzapp.App {
	return &App{
		mux:      cvzruntime.NewServeMux(),
		registry: registry.New(),
	}
}

type App struct {
	mux      cvzruntime.IServeMux
	registry registry.IRegistry
}

func (app *App) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	if err := app.registry.DiscoveryService(ctx, app.mux); err != nil {
		return err
	}

	// Listen HTTP server (and mux calls to gRPC server endpoint)
	log.Printf("HTTP server listening on %s \n", ":9000")
	return app.mux.Listen(":9000")
}
