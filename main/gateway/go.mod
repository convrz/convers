module github.com/convrz/convers/main/gateway

go 1.23.3

replace github.com/convrz/convers/api => ./../../api

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.24.0
	github.com/convrz/convers/api v0.0.0
	google.golang.org/grpc v1.69.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.35.2-20241127180247-a33202765966.1 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241209162323-e6fa225c2576 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241206012308-a4fef0638583 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)
