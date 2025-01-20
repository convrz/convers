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

// Package app implements the Greeter service.
package app

import (
	"fmt"
	"net"

	"github.com/convrz/convers/api/services/greeter/v1"
	"github.com/convrz/convers/core/cvzapp"
	"github.com/convrz/convers/pkg/logger"
	"github.com/convrz/convers/x/greeter/v1/internal/controllers"

	"github.com/convrz/convers/core/cvzservice"
)

// Greeter implements GreeterServiceServer.
type Greeter struct {
	*cvzservice.ServiceServer
	srv greeter.GreeterServiceServer
}

// New creates a new Greeter module.
func New(server controllers.IGreeter) cvzapp.Application {
	return &Greeter{
		ServiceServer: cvzservice.NewDefault(),
		srv:           server,
	}
}

// Run implements IGreeter.
func (g *Greeter) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8000))
	if err != nil {
		return err
	}

	// Listen gRPC srv here
	greeter.RegisterGreeterServiceServer(g.AsServer(), g.srv)
	logger.Infof("gRPC srv listening on %s \n", listener.Addr().String())

	return g.AsServer().Serve(listener)
}
