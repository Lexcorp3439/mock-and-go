package grpc

// Helper структура, упрощающая доступ к API тестируемых сервисов
type Helper struct {
	Connection
}

// NewGrpcHelper - конструктор для GrpcHelper
func NewGrpcHelper() *Helper {
	return &Helper{}
}
