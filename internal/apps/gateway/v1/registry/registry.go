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

package registry

import (
	"context"
	"github.com/convrz/convers/core/cvzruntime"
	"github.com/convrz/convers/internal/apps/gateway/v1/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New() *Registry {
	return &Registry{}
}

type IRegistry interface {
	DiscoveryService(ctx context.Context, mux cvzruntime.IServeMux) error
}

type Registry struct{}

func (r *Registry) DiscoveryService(ctx context.Context, mux cvzruntime.IServeMux) error {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := cvzruntime.RegisterService(ctx, mux, ":8000", opts, new(services.Greeter)); err != nil {
		return err
	}

	return nil
}

func getServiceAddress(name string) string {
	return ":8000"
}
