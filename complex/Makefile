LOCAL_BIN:=$(CURDIR)/bin
LOCAL_DB_DSN:="postgres://postgres:postgres@localhost:5440/complex?sslmode=disable"
LOCAL_DB_NAME:=complex
GO_COVER_EXCLUDE:= "(internal/mock/**/*|*_mock.go|*_minimock.go|models_gen.go|generated.go|swagger.go|*.pb.go|.pb.*.go)"

install:
	go get \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		google.golang.org/protobuf/cmd/protoc-gen-go

bin-deps:
	$(info Installing binary dependencies...)

	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.0.1 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.11.2 && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.11.2 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 && \
	GOBIN=$(LOCAL_BIN) go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@v0.3.0 && \
	GOBIN=$(LOCAL_BIN) go install github.com/mitchellh/gox@v1.0.1  && \
	GOBIN=$(LOCAL_BIN) go install golang.org/x/tools/cmd/goimports@v0.1.9 && \
	GOBIN=$(LOCAL_BIN) go install github.com/golang/mock/mockgen@v1.6.0

gen:
	protoc -I ./proto \
	  --plugin=protoc-gen-grpc-gateway=$(LOCAL_BIN)/protoc-gen-grpc-gateway \
	  --go_out ./pkg/api/public --go_opt paths=source_relative \
	  --go-grpc_out ./pkg/api/public --go-grpc_opt paths=source_relative \
	  --grpc-gateway_out ./pkg/api/public --grpc-gateway_opt paths=source_relative \
	  --proto_path=api/public public.proto

	protoc --proto_path=api/complex complex.proto --go-grpc_out=. --go_out=.

run:
	go run main.go

test:
	rm -rf cover.out]
	DATABASE_DSN=$(LOCAL_DB_DSN) go test -race -coverprofile=cover.out ./...

coverage:
	go tool cover -html=cover.out

db-create-migration: NAME=$NAME
db-create-migration:
	$(LOCAL_BIN)/goose -dir db/migrations postgres "$(LOCAL_DB_DSN)" create "${NAME}" sql

db-up:
	$(LOCAL_BIN)/goose -dir db/migrations postgres "$(DATABASE_DSN)"  up

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
	psql -c "drop database if exists not $(LOCAL_DB_NAME) with (FORCE)"
	psql -c "create database $(LOCAL_DB_NAME)"
	make db-up


