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

// Package server implements the Greeter service server.
package server

import (
	"fmt"
	"net"

	"github.com/convrz/convers/api/gen/go/services/greeter/v1"
	"github.com/convrz/convers/pkg/core/cvzapp"
	"github.com/convrz/convers/pkg/core/cvzservice"
	"github.com/convrz/convers/pkg/logger"
	"github.com/convrz/convers/pkg/msgf"
	"github.com/convrz/convers/x/greeter/v1/internal/controllers"
)

var _ cvzapp.Server = (*Greeter)(nil)

// Greeter implements GreeterServiceServer.
type Greeter struct {
	*cvzservice.ServiceServer
	srv greeter.GreeterServiceServer
}

// ListenAndServe implements IGreeter.
func (g *Greeter) ListenAndServe() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		return err
	}

	// Listen gRPC srv here
	greeter.RegisterGreeterServiceServer(g.AsServer(), g.srv)
	logger.Infof(msgf.InfoGrpcServer, listener.Addr().String())

	return g.AsServer().Serve(listener)
}

// New creates a new Greeter module.
func New(srv controllers.IGreeter) cvzapp.Server {
	return &Greeter{
		ServiceServer: cvzservice.NewDefault(),
		srv:           srv,
	}
}
