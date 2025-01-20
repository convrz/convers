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

// Package gateway provides the gateway app.
package gateway

import (
	"context"

	"github.com/convrz/convers/core/cvzapp"
	"github.com/convrz/convers/core/cvzruntime"
	"github.com/convrz/convers/internal/engine/gateway/services/greeter"
	"github.com/convrz/convers/internal/engine/gateway/types"
	"github.com/convrz/convers/internal/engine/gateway/visitor"
	"github.com/convrz/convers/pkg/log"
)

// New creates a new gateway app
func New() cvzapp.Application {
	return &Gateway{
		mux:     cvzruntime.NewServeMux(),
		visitor: visitor.New(),
	}
}

// Gateway represents the gateway app
type Gateway struct {
	mux     cvzruntime.IServeMux
	visitor types.IVisitor
}

func (g *Gateway) visit(ctx context.Context, services ...types.IService) error {
	for _, service := range services {
		if err := service.Accept(ctx, g.mux, g.visitor); err != nil {
			return err
		}
	}

	return nil
}

func (g *Gateway) register(ctx context.Context) error {
	// TODO: Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	services := []types.IService{
		&greeter.Greeter{},
	}

	return g.visit(ctx, services...)
}

// Run the gateway app
func (g *Gateway) Run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := g.register(ctx); err != nil {
		return err
	}

	// Listen HTTP server (and mux calls to gRPC server endpoint)
	log.Infof("HTTP server listening on %s \n", ":9000")
	return g.mux.Listen(":9000")
}
