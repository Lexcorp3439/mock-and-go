package mocker

import (
	"context"
	"google.golang.org/grpc/metadata"
	"os"
)

const (
	MockKey = "mock"
	MockEnv = "stg"
)

func UseMock(ctx context.Context) bool {
	env, ok := os.LookupEnv("env")
	if !ok || env != MockEnv {
		return false
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if values := md.Get(MockKey); !ok || len(values) == 0 {
		return false
	}
	return true
}
