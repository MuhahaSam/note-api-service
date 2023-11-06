package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/MuhahaSam/golangPractice/internal/app/api/note_v1"
	desc "github.com/MuhahaSam/golangPractice/pkg/note_v1"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const (
	httpPort = ":7000"
	grpcPort = ":7002"
)

// App ...
type App struct {
	noteImpl        *note_v1.Implementation
	serviceProvider *serviceProvider

	pathConfig string

	grpcServer *grpc.Server
	mux        *runtime.ServeMux
}

// NewApp ...
func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}
	err := a.initDeps(ctx)

	return a, err
}

// Run ...
func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	err := a.runGRPC(wg)
	if err != nil {
		return err
	}

	err = a.runPublicHTTP(wg)
	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPublicHTTPHandlers,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)
	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.noteImpl = note_v1.NewNoteV1(
		a.serviceProvider.GetNoteService(ctx),
	)

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.StreamInterceptor(
			grpcMiddleware.ChainStreamServer(
				grpcValidator.StreamServerInterceptor(),
			),
		),

		grpc.UnaryInterceptor(
			grpcMiddleware.ChainUnaryServer(
				grpcValidator.UnaryServerInterceptor(),
			),
		),
	)

	desc.RegisterNoteServiceServer(a.grpcServer, a.noteImpl)

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()

	//nolint:staticcheck
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterNoteServiceHandlerFromEndpoint(ctx, a.mux, grpcPort, opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPC(wg *sync.WaitGroup) error {
	grpcAddress := net.JoinHostPort(
		a.serviceProvider.GetConfig().GrpcConf.Host,
		a.serviceProvider.GetConfig().GrpcConf.Port,
	)
	list, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	go func() {
		defer wg.Done()

		if err = a.grpcServer.Serve(list); err != nil {
			log.Fatalf("failed to process gRPC server: %s", err.Error())
		}
	}()

	log.Printf("Run gRPC server on %s port\n", a.serviceProvider.GetConfig().GrpcConf.Port)
	return nil
}

func (a *App) runPublicHTTP(wg *sync.WaitGroup) error {
	httpAddress := net.JoinHostPort(
		a.serviceProvider.GetConfig().HttpConf.Host,
		a.serviceProvider.GetConfig().HttpConf.Port,
	)
	go func() {
		defer wg.Done()

		if err := http.ListenAndServe(httpAddress, a.mux); err != nil {
			log.Fatalf("failed to process muxer: %s", err.Error())
		}
	}()

	log.Printf("Run public http handler on %s port\n", a.serviceProvider.GetConfig().HttpConf.Port)
	return nil
}
