protoc:
    protoc -I ./internal/pkg/proto ./internal/pkg/proto/sso/sso.proto \
		--go_out=./internal/pkg/proto/gen/go \
        --go_opt=paths=source_relative \
        --go-grpc_out=./internal/pkg/proto/gen/go \
        --go-grpc_opt=paths=source_relative