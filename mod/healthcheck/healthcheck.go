package healthcheck

import (
	"fmt"
	"log"
	"net"

	"github.com/makegalxy/galxy/mod/healthcheck/internal/handlers"
	"github.com/makegalxy/galxy/pkg/proto/healthcheck"
	"google.golang.org/grpc"
)

type _Greeter struct {
	grpcServer *grpc.Server
	service    healthcheck.GreeterServer
}

// NewGreeter creates a new Greeter service.
func NewGreeter() *_Greeter {
	return &_Greeter{
		grpcServer: grpc.NewServer(),
		service:    handlers.NewGreeter(),
	}
}

// Serve implements IGreeter.
func (g *_Greeter) Serve() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8000))
	if err != nil {
		return err
	}

	// Start gRPC server here
	healthcheck.RegisterGreeterServer(g.grpcServer, g.service)
	log.Printf("gRPC server listening on %s \n", listener.Addr().String())
	return g.grpcServer.Serve(listener)
}
