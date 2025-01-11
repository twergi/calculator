package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twergi/calculator/internal/app/components/grpcctrl"
	"github.com/twergi/calculator/internal/app/components/httpctrl"
	"github.com/twergi/calculator/internal/app/repository"
	"github.com/twergi/calculator/internal/app/usecases/calculator"
	"github.com/twergi/calculator/internal/config"
	desc "github.com/twergi/calculator/internal/proto/gen/go/service"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	log.Println("err:", run().Error())
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	config, err := config.NewFromFile()
	if err != nil {
		return err
	}

	pool, err := pgxpool.New(
		ctx,
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.DB.Host,
			config.DB.Port,
			config.DB.User,
			config.DB.Password,
			config.DB.Database,
			config.DB.Sslmode,
		),
	)
	if err != nil {
		return err
	}

	repo := repository.New(pool)
	calcUsecase := calculator.New(repo)
	httpServer := httpctrl.New("127.0.0.1:8000", calcUsecase)
	grpcImpl := grpcctrl.New(calcUsecase)
	grpcListener, err := (&net.ListenConfig{}).Listen(ctx, "tcp", "localhost:8001")
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	desc.RegisterCalculatorServer(server, grpcImpl)

	errGr, errCtx := errgroup.WithContext(ctx)
	errGr.Go(func() error {
		return httpServer.Run(ctx)
	})
	errGr.Go(func() error {
		<-errCtx.Done()
		return httpServer.Stop(errCtx)
	})

	errGr.Go(func() error {
		slog.InfoContext(ctx, "starting grpc server", "addr", grpcListener.Addr())
		return server.Serve(grpcListener)
	})
	errGr.Go(func() error {
		<-errCtx.Done()
		slog.InfoContext(ctx, "stopping grpc server")
		server.GracefulStop()
		return nil
	})

	return errGr.Wait()
}
