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

package app

import (
	"context"
	gw "github.com/convrz/convers/api/greeter/v1" // Update
	"github.com/convrz/convers/core/apps"
	"github.com/convrz/convers/core/servers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func New() apps.App {
	return &App{
		server: servers.New(":9000"),
	}
}

type App struct {
	server servers.IServer
}

func (app *App) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := app.server.Register(ctx, gw.RegisterGreeterHandlerFromEndpoint, ":8000", opts...); err != nil {
		return err
	}

	log.Printf("HTTP server listening on %s \n", ":9000")

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return app.server.Start()
}
