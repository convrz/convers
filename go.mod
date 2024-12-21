module github.com/convrz/convers

go 1.23.3

replace github.com/convrz/convers/api => ./api

require (
	github.com/convrz/convers/api v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.6.0
	github.com/jackc/pgx/v5 v5.7.1
	google.golang.org/protobuf v1.35.2
)

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.35.2-20241127180247-a33202765966.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
