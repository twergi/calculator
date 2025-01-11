package calculator

import "context"

type Repository interface {
	SaveResult(ctx context.Context, result int64) error
	GetLastResult(ctx context.Context) (int64, error)
}

type Usecase struct {
	repo Repository
}

func New(repo Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}
