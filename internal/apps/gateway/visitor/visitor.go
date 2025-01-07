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

package visitor

import (
	"context"

	"github.com/convrz/convers/v1/core/cvzruntime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New() IVisitor {
	return &Visitor{}
}

var _ IVisitor = (*Visitor)(nil)

type Visitor struct{}

func (v *Visitor) VisitGreeterService(ctx context.Context, mux cvzruntime.IServeMux, service IService) error {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	greeterAddr := ":8000"
	if err := cvzruntime.RegisterService(ctx, mux, service, greeterAddr, opts); err != nil {
		return err
	}

	return nil
}
