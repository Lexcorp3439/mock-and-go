package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	c "heisenbug/complex/internal/app/complex"
	"heisenbug/complex/internal/app/public"
	"heisenbug/complex/internal/pkg/repository"
	"heisenbug/complex/internal/pkg/store"
	desc2 "heisenbug/complex/pkg/api/complex"
	desc1 "heisenbug/complex/pkg/api/public"
)

func main() {
	db, err := store.ConnectToPostgres()
	if err != nil {
		panic(err)
	}

	storage := store.NewStorage(db)
	repo := repository.NewTemplate(storage)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	complexServ := c.NewComplexService(repo)
	publicServ := public.NewPublicService(repo)

	s := grpc.NewServer()
	desc1.RegisterPublicServer(s, publicServ)
	desc2.RegisterComplexServer(s, complexServ)

	reflection.Register(s)

	go func() {
		if err = s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = desc1.RegisterPublicHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
