LOCAL_BIN=$(CURDIR)/bin

install-deps:
	@GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	@GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

get-deps:
	@go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	@go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-auth-api

generate-auth-api:
	@mkdir -p pkg/chat_v1
	@protoc \
		--proto_path=api \
		--go_out=paths=source_relative:pkg \
		--go-grpc_out=paths=source_relative:pkg \
		api/chat_v1/chat.proto