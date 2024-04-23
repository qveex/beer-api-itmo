package app

import (
	"log/slog"
	"time"

	grpcapp "grpc-service-ref/internal/app/grpc"
	"grpc-service-ref/internal/services/auth"
	"grpc-service-ref/internal/storage/postgres"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	databaseURL string,
	tokenTTL time.Duration,
) *App {
	storage, err := postgres.New(databaseURL)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
