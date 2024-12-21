module github.com/convrz/convers/x/greeter

go 1.23.3

replace (
	github.com/convrz/convers/api => ./../../api
	github.com/convrz/convers/core => ./../../core
)

require (
	github.com/convrz/convers/api v0.0.0-00010101000000-000000000000
	github.com/convrz/convers/core v0.0.0-20241215153253-12c2143a5f64
	google.golang.org/grpc v1.69.0
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.35.2-20241127180247-a33202765966.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.24.0 // indirect
	golang.org/x/net v0.32.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241209162323-e6fa225c2576 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241209162323-e6fa225c2576 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)
