package main

import (
	"github.com/makegalxy/galxy/mod/gateway"
	"google.golang.org/grpc/grpclog"
)

func main() {
	if err := gateway.Serve(); err != nil {
		grpclog.Fatal(err)
	}
}
