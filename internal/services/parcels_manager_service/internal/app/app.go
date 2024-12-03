package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/subliker/track-parcel-service/internal/pkg/broker/rabbitmq"
	"github.com/subliker/track-parcel-service/internal/pkg/logger"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/config"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/server/rest/api"
	"github.com/subliker/track-parcel-service/internal/services/parcels_manager_service/internal/store/parcel"
	"google.golang.org/grpc"
)

// App is interface to control app.
// Closing app closes all app services.
type App interface {
	// Run starts app and stops all internal services when ctx done.
	Run(context.Context) error
}

type app struct {
	parcelServer *grpc.Server
	grpcAddress  string

	apiServer *api.Server

	store parcel.ManagerStore

	broker rabbitmq.Broker

	logger logger.Logger
}

// AppOptions is struct for building app arguments
type AppOptions struct {
	Config       config.Config
	Store        parcel.ManagerStore
	ParcelServer *grpc.Server
	APIServer    *api.Server
	Broker       rabbitmq.Broker
}

// New creates new instance of app
func New(logger logger.Logger, opts AppOptions) App {
	var a app

	// setting grpc address
	a.grpcAddress = fmt.Sprintf(":%d", opts.Config.GRPC.Port)

	// setting logger
	a.logger = logger.WithFields("layer", "app")

	// setting store
	a.store = opts.Store

	// setting parcel server
	a.parcelServer = opts.ParcelServer

	// setting rest api service
	a.apiServer = opts.APIServer

	// setting broker
	a.broker = opts.Broker
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
		if err := a.parcelServer.Serve(lis); err != nil {
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

	wg.Add(1)
	go func() {
		defer wg.Done()
		// starting api server
		if err := a.apiServer.Run(); err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				errMsg := fmt.Errorf("running rest api server error: %s", err)
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
	// closing rest api server
	if err := a.apiServer.Close(); err != nil {
		a.logger.Warnf("rest api server closing ended with error: %s", err)
	}
	// wait until services will be stopped
	wg.Wait()

	// stop broker
	if err := a.broker.Close(); err != nil {
		a.logger.Warn(err)
	}

	// close store
	a.store.Close()

	a.logger.Info("app was gracefully shutdowned :)")
	return nil
}
