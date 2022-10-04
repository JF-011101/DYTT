package grpc

import (
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// RecoveryInterceptor return Unknow when panic occurs.
func RecoveryInterceptor() grpcrecovery.Option {
	return grpcrecovery.WithRecoveryHandler(func(p interface{}) (err error) {
		return grpc.Errorf(codes.Unknown, "panic triggered: %v", p)
	})
}
