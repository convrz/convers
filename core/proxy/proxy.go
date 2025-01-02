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

package proxy

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"net/http"
)

func New(addr string, opts ...runtime.ServeMuxOption) IProxy {
	return &Proxy{
		addr: addr,
		mux:  runtime.NewServeMux(opts...),
	}
}

type RegisterHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

type IProxy interface {
	Register(ctx context.Context, handler RegisterHandler, serviceHost string, opts ...grpc.DialOption) error
	Start() error
}

type Proxy struct {
	mux  *runtime.ServeMux
	addr string
}

func (proxy *Proxy) Start() error {
	return http.ListenAndServe(proxy.addr, proxy.mux)
}

func (proxy *Proxy) Register(ctx context.Context, handler RegisterHandler, serviceHost string, opts ...grpc.DialOption) error {
	return handler(ctx, proxy.mux, serviceHost, opts)
}
