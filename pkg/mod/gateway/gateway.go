package gateway

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "github.com/makegalxy/galxy/pkg/proto/healthcheck" // Update
)

func Serve() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, "localhost:8000", opts)
	if err != nil {
		return err
	}

	log.Printf("HTTP server listening on %s \n", ":9000")

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":9000", mux)
}
