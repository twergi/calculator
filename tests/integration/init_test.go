package integration

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/twergi/calculator/internal/app/components/grpcctrl"
	"github.com/twergi/calculator/internal/app/repository"
	"github.com/twergi/calculator/internal/app/usecases/calculator"
)

var (
	ctx    context.Context
	cancel context.CancelFunc

	pool *pgxpool.Pool
)

func init() {
	ctx, cancel = context.WithCancel(context.Background())

	var err error
	pool, err = pgxpool.New(ctx, "host=localhost port=5432 user=postgres password=postgres dbname=calculator sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func newGRPC() *grpcctrl.Implementation {
	return grpcctrl.New(calculator.New(repository.New(pool)))
}

func truncate() {
	_, err := pool.Exec(context.Background(), `
TRUNCATE TABLE results CASCADE;
`)
	if err != nil {
		panic(err)
	}
}
