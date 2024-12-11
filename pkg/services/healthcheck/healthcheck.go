/*
Copyright 2024 The Galxy Authors.

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

package healthcheck

import (
	"fmt"
	"log"
	"net"

	"github.com/makegalxy/galxy/pkg/proto/healthcheck"
	"github.com/makegalxy/galxy/pkg/services/healthcheck/internal/handlers"
	"google.golang.org/grpc"
)

type _Greeter struct {
	grpcServer *grpc.Server
	service    healthcheck.GreeterServer
}

// NewGreeter creates a new Greeter service.
func NewGreeter() *_Greeter {
	return &_Greeter{
		grpcServer: grpc.NewServer(),
		service:    handlers.NewGreeter(),
	}
}

// Run implements IGreeter.
func (g *_Greeter) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8000))
	if err != nil {
		return err
	}

	// Start gRPC server here
	healthcheck.RegisterGreeterServer(g.grpcServer, g.service)
	log.Printf("gRPC server listening on %s \n", listener.Addr().String())
	return g.grpcServer.Serve(listener)
}
