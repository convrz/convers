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

package services

import (
	"google.golang.org/grpc"
)

var _ IServiceRegistrar = (*ServiceRegistrar)(nil)

type IServiceRegistrar interface {
	Run() error
}

func New(opts ...grpc.ServerOption) *ServiceRegistrar {
	return &ServiceRegistrar{
		Server: grpc.NewServer(opts...),
	}
}

func NewDefault() *ServiceRegistrar {
	return New()
}

type ServiceRegistrar struct {
	*grpc.Server
}

func (s *ServiceRegistrar) Run() error {
	panic("unimplemented")
}
