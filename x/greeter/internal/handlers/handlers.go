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

package handlers

import (
	"context"

	"github.com/convrz/convers/api/services/greeter/v1"
)

// Greeter is the module for Greeter.
type Greeter struct {
	greeter.UnimplementedGreeterServer
}

// NewGreeter creates a new Greeter module.
func NewGreeter() greeter.GreeterServer {
	return &Greeter{}
}

// SayHello implements GreeterServer.
func (g *Greeter) SayHello(_ context.Context, msg *greeter.HelloRequest) (*greeter.HelloReply, error) {
	name := msg.GetName()
	return &greeter.HelloReply{
		Message: "Reply " + name,
	}, nil
}
