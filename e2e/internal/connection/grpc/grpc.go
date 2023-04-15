package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ConnectionFactory is an interface to implement your own ConnectionFactory
type ConnectionFactory interface {
	GrpcUnaryConnectionFactory(testCtx interface{}, host string) *grpc.ClientConn
}

// Connection provides function to get *grpc.ClientConn.
type Connection struct{}

// GrpcUnaryConnectionFactory returns *grpc.ClientConn
func (g *Connection) GrpcUnaryConnectionFactory(host string) *grpc.ClientConn {
	conn, _ := grpc.Dial(
		host,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	return conn
}
