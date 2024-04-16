package di

import (
	ssov1 "github.com/Krutov777/protos/gen/go/sso"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func (sp ServiceProvider) GetAuthClient() *ssov1.AuthClient {
	if sp.authClient != nil {
		return sp.authClient
	}

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(3)),
		grpcretry.WithPerRetryTimeout(time.Duration(5) * time.Second),
	}

	conn, err := grpc.NewClient(
		":44044",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(grpcretry.UnaryClientInterceptor(retryOpts...)),
	)

	if err != nil {
		panic(err)
	}

	client := ssov1.NewAuthClient(conn)
	sp.authClient = &client

	return sp.authClient
}
