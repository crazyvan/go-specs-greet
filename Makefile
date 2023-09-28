.PHONY: test
test:
	go test ./...

.PHONY: protoc
protoc:
	cd adapters/grpcserver && \
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    greet.proto