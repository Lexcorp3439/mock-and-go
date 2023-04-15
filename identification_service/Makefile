LOCAL_BIN:=$(CURDIR)/bin

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	go get google.golang.org/grpc
	brew install protobuf
	brew install clang-format
	brew install grpcurl
	export PATH=$PATH:$(go env GOPATH)/bin

bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0

gen:
	protoc --proto_path=api/identification identification.proto --go-grpc_out=. --go_out=.

gen_clients:
	protoc --proto_path=external users.proto --go-grpc_out=. --go_out=.

run:
	go run main.go

test:
	rm -rf cover.out
	go test -race -coverprofile=cover.out ./...
	go tool cover -html=cover.out

mocks:
	rm -rf ./internal/pkg/mocks
	$(LOCAL_BIN)/mockgen -package=users_mock -destination=./pkg/mocks/users_mock/users_mock.go -source=./internal/pb/users/users_grpc.pb.go UsersClient
	$(LOCAL_BIN)/mockgen -package=external_mock -destination=./pkg/mocks/external_mock/external_mock.go -source=./internal/pkg/external/client.go Client


