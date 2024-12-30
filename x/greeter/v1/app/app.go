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
	"fmt"
	"github.com/convrz/convers/api/greeter/v1"
	"github.com/convrz/convers/core/apps"
	"github.com/convrz/convers/x/greeter/v1/internal/controllers"
	"log"
	"net"

	"github.com/convrz/convers/core/services"
)

type Greeter struct {
	*services.ServiceRegistrar
	srv greeter.GreeterServer
}

// New creates a new Greeter module.
func New(service *services.ServiceRegistrar, server controllers.IGreeter) apps.App {
	return &Greeter{
		ServiceRegistrar: service,
		srv:              server,
	}
}

// Run implements IGreeter.
func (g *Greeter) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8000))
	if err != nil {
		return err
	}

	// Start gRPC srv here
	greeter.RegisterGreeterServer(g.ServiceRegistrar, g.srv)
	log.Printf("gRPC srv listening on %s \n", listener.Addr().String())
	return g.Serve(listener)
}
