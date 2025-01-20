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

// Package cvzservice provides a service registrar.
package cvzservice

import (
	"google.golang.org/grpc"
)

var _ IServiceServer = (*ServiceServer)(nil)

// IServiceServer is a service registrar.
type IServiceServer interface {
	AsServer() *grpc.Server
	Run() error
}

// New returns a new service registrar.
func New(opts ...grpc.ServerOption) *ServiceServer {
	return &ServiceServer{
		server: grpc.NewServer(opts...),
	}
}

// NewDefault returns a new service registrar with default options.
func NewDefault() *ServiceServer {
	return New()
}

// ServiceServer is a gRPC server that registers services.
type ServiceServer struct {
	server *grpc.Server
}

// AsServer returns the underlying gRPC server.
func (s *ServiceServer) AsServer() *grpc.Server {
	return s.server
}

// Run starts the service registrar.
func (s *ServiceServer) Run() error {
	panic("unimplemented")
}
