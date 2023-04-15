package provider

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/ozontech/allure-go/pkg/allure"
)

type ProviderT interface {
	Step(step *allure.Step)
	Logf(format string, args ...interface{})
	Name() string
}

type testCtxKey struct{}

func BackgroundWithProvider(t ProviderT) context.Context {
	return WithProvider(context.Background(), t)
}

func WithProvider(parentCtx context.Context, t ProviderT) context.Context {
	return context.WithValue(parentCtx, testCtxKey{}, t)
}

func GetProviderT(ctx context.Context) (ProviderT, error) {
	testCtx := ctx.Value(testCtxKey{})
	if testCtx == nil {
		errMsg := "TestContext did not passed with connection.WithProviderT method"
		return nil, errors.New(errMsg)
	}
	if reflect.ValueOf(testCtx).Kind() != reflect.Ptr {
		errMsg := fmt.Sprintf("Value is not a pointer (%v)", testCtx)
		return nil, errors.New(errMsg)
	}
	if p, ok := testCtx.(ProviderT); ok {
		return p, nil
	}
	errMsg := fmt.Sprintf("Wrong pointer type. Expected: (provider.T). Actual: (%v)", testCtx)
	return nil, errors.New(errMsg)
}
