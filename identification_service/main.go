package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	impl "heisenbug/identification/internal/app"
	compl "heisenbug/identification/internal/pkg/complex"
	"heisenbug/identification/internal/pkg/external"
	"heisenbug/identification/internal/pkg/users"
	desc "heisenbug/identification/pkg/api/identification"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	usersClient, usersConn, err := users.NewClient()
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer usersConn.Close()

	externalClient := external.NewClient()

	complexClient := compl.NewClient()

	serv := impl.NewIdentificationService(usersClient, externalClient, complexClient)

	s := grpc.NewServer()
	desc.RegisterIdentificationServer(s, serv)
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
