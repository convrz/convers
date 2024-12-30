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

package controllers

import (
	"context"
	greeterBiz "github.com/convrz/convers/api/biz/greeter/v1"
	"github.com/convrz/convers/api/greeter/v1"
	"github.com/convrz/convers/pkg/copier"
	"github.com/convrz/convers/x/greeter/v1/internal/biz"
	"github.com/convrz/convers/x/greeter/v1/internal/repos"
)

type IGreeter interface {
	greeter.GreeterServer
}

// Greeter is the module for Greeter.
type Greeter struct {
	greeter.UnimplementedGreeterServer
	repos repos.IDB
	biz   greeterBiz.GreeterServer
}

// New creates a new Greeter module.
func New(biz biz.IGreeter) IGreeter {
	return &Greeter{
		biz: biz,
	}
}

// SayHello implements GreeterServer.
func (g *Greeter) SayHello(ctx context.Context, msg *greeter.HelloRequest) (*greeter.HelloReply, error) {
	var SayHelloReq greeterBiz.HelloRequest
	if err := copier.CopyMsg(msg, &SayHelloReq); err != nil {
		return nil, err
	}

	sayHelloResp, err := g.biz.SayHello(ctx, &SayHelloReq)
	if err != nil {
		return nil, err
	}

	var resp greeter.HelloReply
	if err := copier.CopyMsg(sayHelloResp, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
