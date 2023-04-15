package connection

import (
	"e2e/internal/connection/grpc"
	"e2e/internal/pb/complex"
	"e2e/internal/pb/identification"
	"e2e/internal/pb/users"
	"github.com/walkerus/go-wiremock"
)

func NewIdentificationClient(h *grpc.Helper) identification.IdentificationClient {
	return identification.NewIdentificationClient(h.GrpcUnaryConnectionFactory(IdentificationHost))
}

func NewUsersClient(h *grpc.Helper) users.UsersClient {
	return users.NewUsersClient(h.GrpcUnaryConnectionFactory(UsersHost))
}

func NewComplexClient(h *grpc.Helper) complex.ComplexClient {
	return complex.NewComplexClient(h.GrpcUnaryConnectionFactory(ComplexHost))
}

func NewWiremockClient() *wiremock.Client {
	return wiremock.NewClient(WiremockHost)
}

//func NewComplexClient(h *http.Helper) (*complex.Client)
