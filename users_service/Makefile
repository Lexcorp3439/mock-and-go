LOCAL_BIN:=$(CURDIR)/bin
LOCAL_DB_DSN:="postgres://postgres:postgres@localhost:5441/users?sslmode=disable"
LOCAL_DB_NAME:=users
GO_COVER_EXCLUDE:= "(internal/mock/**/*|*_mock.go|*_minimock.go|models_gen.go|generated.go|swagger.go|*.pb.go|.pb.*.go)"

install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	go get google.golang.org/grpc
	brew install protobuf
	brew install clang-format
	brew install grpcurl
	export PATH=$PATH:$(go env GOPATH)/bin

bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.0.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0

gen:
	protoc --proto_path=api/users users.proto --go-grpc_out=. --go_out=.

run:
	go run main.go

db-create-migration: NAME=$NAME
db-create-migration:
	$(LOCAL_BIN)/goose -dir db/migrations postgres "$(LOCAL_DB_DSN)" create "${NAME}" sql

db-up:
	$(LOCAL_BIN)/goose -dir db/migrations postgres "$(DATABASE_DSN)" up

db-up-local:
	$(LOCAL_BIN)/goose -dir db/migrations postgres "$(LOCAL_DB_DSN)" up

db-status:
	$(LOCAL_BIN)/goose -dir db/migrations postgres "$(LOCAL_DB_DSN)" status

db-migrate:
	make db-up
	make db-gen-structure

db-migrate-down:
	$(LOCAL_BIN)/goose -dir db/migrations postgres $(LOCAL_DB_DSN) down

db-reset:
	psql -c "drop database if exists $(LOCAL_DB_NAME) with (FORCE)"
	psql -c "create database $(LOCAL_DB_NAME)"
	make db-up

test:
	rm -rf cover.out
	DATABASE_DSN=$(LOCAL_DB_DSN) go test -race -coverprofile=cover.out ./internal...
	go tool cover -html=cover.out

