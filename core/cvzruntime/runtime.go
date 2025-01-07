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

package cvzruntime

import (
	"context"
	"google.golang.org/grpc"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type GrpcService interface {
	Register(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
}

type IServeMux interface {
	Listen(addr string) error
	AsRuntimeMux() *runtime.ServeMux
}

func NewServeMux(opts ...runtime.ServeMuxOption) IServeMux {
	return &ServeMux{
		ServeMux: runtime.NewServeMux(opts...),
	}
}

type ServeMux struct {
	*runtime.ServeMux
}

func (mux *ServeMux) Listen(addr string) error {
	return http.ListenAndServe(addr, mux)
}

func (mux *ServeMux) AsRuntimeMux() *runtime.ServeMux {
	return mux.ServeMux
}
