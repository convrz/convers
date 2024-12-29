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
	"github.com/convrz/convers/api/services/greeter/v1"
	"log"
	"net"

	"github.com/convrz/convers/core/containers"
	"github.com/convrz/convers/core/servers"
)

type Greeter struct {
	*servers.Server
	service greeter.GreeterServer
}

// New creates a new Greeter module.
func New(server *servers.Server, service greeter.GreeterServer) containers.IApp {
	return &Greeter{
		Server:  server,
		service: service,
	}
}

// Run implements IGreeter.
func (g *Greeter) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8000))
	if err != nil {
		return err
	}

	// Start gRPC server here
	greeter.RegisterGreeterServer(g, g.service)
	log.Printf("gRPC server listening on %s \n", listener.Addr().String())
	return g.Serve(listener)
}
