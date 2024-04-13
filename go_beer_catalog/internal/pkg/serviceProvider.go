package pkg

import (
	"fmt"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"main/internal/domain"
	"main/internal/env"
	"main/internal/repository"
	pb "main/pkg/api"
	"time"
)

type ServiceProvider struct {
	db             *gorm.DB
	authClient     *pb.CatalogClient
	catalogService *domain.CatalogService
	catalogRepo    *repository.CatalogRepository
	authService    *domain.AuthService
	authRepo       *repository.AuthRepository
}

var Provider = &ServiceProvider{}

func (sp *ServiceProvider) GetDb() *gorm.DB {
	if sp.db != nil {
		return sp.db
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		env.GetEnv(dbUserEnv, ""),
		env.GetEnv(dbPassEnv, ""),
		env.GetEnv(dbAddrEnv, ""),
		env.GetEnv(dbPortEnv, ""),
		env.GetEnv(dbNameEnv, ""),
	)

	var err error
	sp.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return sp.db
}

func (sp *ServiceProvider) GetCatalogService() *domain.CatalogService {
	if sp.catalogService != nil {
		return sp.catalogService
	}

	sp.catalogService = domain.NewCatalogService(sp.GetCatalogRepository())

	return sp.catalogService
}

func (sp ServiceProvider) GetCatalogRepository() *repository.CatalogRepository {
	if sp.catalogRepo != nil {
		return sp.catalogRepo
	}

	sp.catalogRepo = repository.NewCatalogRepository(sp.GetDb())

	return sp.catalogRepo
}

func (sp *ServiceProvider) GetAuthService() *domain.AuthService {
	if sp.catalogService != nil {
		return sp.authService
	}

	sp.authService = domain.NewAuthService(sp.GetAuthRepository())

	return sp.authService
}

func (sp ServiceProvider) GetAuthRepository() *repository.AuthRepository {
	if sp.authRepo != nil {
		return sp.authRepo
	}

	sp.authRepo = repository.NewAuthRepository(sp.GetAuthClient())

	return sp.authRepo
}

func (sp ServiceProvider) GetAuthClient() *pb.CatalogClient {
	if sp.authClient != nil {
		return sp.authClient
	}

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(3)),
		grpcretry.WithPerRetryTimeout(time.Duration(5) * time.Second),
	}

	conn, err := grpc.NewClient(
		":8081",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(grpcretry.UnaryClientInterceptor(retryOpts...)),
	)

	if err != nil {
		panic(err)
	}

	client := pb.NewCatalogClient(conn) // TODO Auth
	sp.authClient = &client

	return sp.authClient
}

const (
	dbUserEnv = "DB_USER"
	dbPassEnv = "DB_PASS"
	dbAddrEnv = "DB_ADDR"
	dbPortEnv = "DB_PORT"
	dbNameEnv = "DB_NAME"
)
