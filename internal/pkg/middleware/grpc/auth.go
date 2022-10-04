package grpc

import (
	"context"
	"errors"

	grpcauth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type TokenInfo struct {
	ID    string
	Roles []string
}

// AuthInterceptor validate the token with authorization as the header in the form of 'bearer token'
func AuthInterceptor(ctx context.Context) (context.Context, error) {
	token, err := grpcauth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, " %v", err)
	}

	newCtx := context.WithValue(ctx, tokenInfo.ID, tokenInfo)
	return newCtx, nil
}

func parseToken(token string) (TokenInfo, error) {
	var tokenInfo TokenInfo
	if token == "dytt.grpc.auth.token" {
		tokenInfo.ID = "1"
		tokenInfo.Roles = []string{"admin"}
		return tokenInfo, nil
	}
	return tokenInfo, errors.New("Token无效: bearer " + token)
}
