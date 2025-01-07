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

package base

import (
	"context"

	"github.com/convrz/convers/v1/core/cvzruntime"
	"github.com/convrz/convers/v1/internal/apps/gateway/visitor"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var _ visitor.IService = (*Service)(nil)

type Service struct{}

func (s *Service) Register(_ context.Context, _ *runtime.ServeMux, _ string, _ []grpc.DialOption) error {
	panic("unimplemented")
}

func (s *Service) Accept(_ context.Context, _ cvzruntime.IServeMux, _ visitor.IVisitor) error {
	panic("unimplemented")
}
