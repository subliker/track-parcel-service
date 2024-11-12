package app

import (
	"fmt"
	"net"

	"github.com/subliker/track-parcel-service/internal/pkg/logger/zap"
	ssov1 "github.com/subliker/track-parcel-service/internal/pkg/proto/gen/go/sso"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/grpc/auth"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/store/pgstore"

	"google.golang.org/grpc"
)

type App struct {
	ssoServer *grpc.Server
}

func New(cfg config.Config) *App {
	var a App

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		zap.Logger.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	// up store
	store, err := pgstore.New(cfg.DB)
	if err != nil {
		zap.Logger.Fatal(err)
	}

	ssov1.RegisterAuthServer(grpcServer, auth.New(store))

	fmt.Println(lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		zap.Logger.Fatal(err)
	}

	a.ssoServer = grpcServer
	return &a
}
