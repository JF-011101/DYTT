package grpc

import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// ZapInterceptor return an instance of zap.Logger
func ZapInterceptor() *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  "./tmp/grpc.log",
		MaxSize:   1024, //MB
		LocalTime: true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		zap.NewAtomicLevel(),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	grpc_zap.ReplaceGrpcLogger(logger)
	return logger
}
