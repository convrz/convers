package main

import (
	"github.com/makegalxy/galxy/mod/healthcheck"
	"google.golang.org/grpc/grpclog"
)

func main() {
	server := healthcheck.NewGreeter()
	if err := server.Serve(); err != nil {
		grpclog.Fatal(err)
	}
}
