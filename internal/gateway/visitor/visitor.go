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

// Package visitor provides an implementation of the visitor pattern.
package visitor

import (
	"context"

	cvzruntime "github.com/convrz/convers/pkg/core/cvzruntime"

	"github.com/convrz/convers/internal/gateway/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// New returns a new visitor.
func New() types.IVisitor {
	return &Visitor{}
}

var _ types.IVisitor = (*Visitor)(nil)

// Visitor represents the visitor interface.
type Visitor struct{}

// VisitGreeterService visits the greeter service.
func (v *Visitor) VisitGreeterService(ctx context.Context, mux cvzruntime.IServeMux, service types.IService) error {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	greeterAddr := ":8000"
	if err := cvzruntime.RegisterService(ctx, mux, service, greeterAddr, opts); err != nil {
		return err
	}

	return nil
}
