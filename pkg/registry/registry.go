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

// Package server provides the server for the service registry.
package server

import (
	"context"

	"github.com/convrz/convers/api/gen/go/utils/srx/v1"
)

var _ srx.ServiceRegistryServiceServer = (*ServiceRegistry)(nil)

// ServiceRegistry implements the ServiceRegistryService.
type ServiceRegistry struct {
	srx.UnimplementedServiceRegistryServiceServer
}

// RegisterService implements the RegisterService method of the ServiceRegistryService.
func (srx *ServiceRegistry) RegisterService(_ context.Context, _ *srx.RegisterServiceRequest) (*srx.RegisterServiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

// DiscoverService implements the DiscoverService method of the ServiceRegistryService.
func (srx *ServiceRegistry) DiscoverService(_ context.Context, _ *srx.DiscoverServiceRequest) (*srx.DiscoverServiceResponse, error) {
	//TODO implement me
	panic("implement me")
}

// Heartbeat implements the Heartbeat method of the ServiceRegistryService.
func (srx *ServiceRegistry) Heartbeat(_ context.Context, _ *srx.HeartbeatRequest) (*srx.HeartbeatResponse, error) {
	//TODO implement me
	panic("implement me")
}
