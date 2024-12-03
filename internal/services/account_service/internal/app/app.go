package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/account_service/internal/store"
	"google.golang.org/grpc"
)

// App is interface to control app.
// Closing app closes all app services.
type App interface {
	// Run starts app and stops all internal services when ctx done.
	Run(context.Context) error
}

type app struct {
	accountServer *grpc.Server
	grpcAddress   string

	store store.Store

	logger logger.Logger
}

// New creates new instance of app
func New(cfg config.Config,
	logger logger.Logger,
	store store.Store,
	accountServer *grpc.Server) App {
	var a app

	// setting grpc address
	a.grpcAddress = fmt.Sprintf(":%d", cfg.GRPC.Port)

	// setting store
	a.store = store

	// setting account server
	a.accountServer = accountServer

	// setting logger
	a.logger = logger.WithFields("layer", "app")
	return &a
}

func (a *app) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// receive sys signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// creating new new listener
	lis, err := net.Listen("tcp", a.grpcAddress)
	if err != nil {
		a.logger.Fatal(err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		// starting serving server
		a.logger.Infof("starting grpc server at port %d...", a.grpcAddress)
		if err := a.accountServer.Serve(lis); err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				errMsg := fmt.Errorf("serving grpc server error: %s", err)
				a.logger.Error(errMsg)
				cancel()
			}
		}
	}()
	a.logger.Info("app running...")

	// wait until signal will come or context will end
	select {
	case <-quit:
		a.logger.Info("shutdown signal received")
	case <-ctx.Done():
		a.logger.Info("context canceled")
	}

	a.logger.Info("stopping all services")

	// stop services
	cancel()
	// closing net listener
	if err := lis.Close(); err != nil {
		a.logger.Warnf("net listener closing ended with error: %s", err)
	}
	// wait until services will be stopped
	wg.Wait()

	// close store
	a.store.Close()

	a.logger.Info("app was gracefully shutdowned :)")
	return nil
}
