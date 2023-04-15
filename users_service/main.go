package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"heisenbug/users/internal/app/users"
	"heisenbug/users/internal/pkg/repository"
	"heisenbug/users/internal/pkg/store"
	desc "heisenbug/users/pkg/api/users"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	db, err := store.ConnectToPostgres()
	if err != nil {
		panic(err)
	}

	storage := store.NewStorage(db)
	repo := repository.NewUsersRepository(storage)

	serv := users.NewUsersService(repo)

	s := grpc.NewServer()
	desc.RegisterUsersServer(s, serv)
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
