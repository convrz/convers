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

package servers

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net/http"
)

func New(addr string) IServer {
	return &Server{
		addr: addr,
		mux:  runtime.NewServeMux(),
	}
}

type RegisterHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

type IServer interface {
	Register(ctx context.Context, handler RegisterHandler, serviceHost string, opts ...grpc.DialOption) error
	Start() error
}

type Server struct {
	mux  *runtime.ServeMux
	addr string
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.addr, s.mux)
}

func (s *Server) Register(ctx context.Context, handler RegisterHandler, serviceHost string, opts ...grpc.DialOption) error {
	return handler(ctx, s.mux, serviceHost, opts)
}
