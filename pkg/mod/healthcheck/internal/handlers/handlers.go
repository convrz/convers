package handlers

import (
	"context"

	"github.com/makegalxy/galxy/pkg/proto/healthcheck"
)

// Greeter is the service for Greeter.
type Greeter struct {
	healthcheck.UnimplementedGreeterServer
}

// NewGreeter creates a new Greeter service.
func NewGreeter() healthcheck.GreeterServer {
	return &Greeter{}
}

// SayHello implements GreeterServer.
func (g *Greeter) SayHello(_ context.Context, msg *healthcheck.HelloRequest) (*healthcheck.HelloReply, error) {
	name := msg.GetName()
	return &healthcheck.HelloReply{
		Message: "Reply " + name,
	}, nil
}
