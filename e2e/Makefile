LOCAL_BIN:=$(CURDIR)/bin

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	go get google.golang.org/grpc
	brew install protobuf
	brew install clang-format
	brew install grpcurl
	export PATH=$PATH:$(go env GOPATH)/bin

gen_clients:
	protoc --proto_path=external identification.proto --go-grpc_out=. --go_out=.
	protoc --proto_path=external users.proto --go-grpc_out=. --go_out=.
	protoc --proto_path=external complex.proto --go-grpc_out=. --go_out=.

test:
	go test ./...


