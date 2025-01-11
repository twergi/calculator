package calculator

import "context"

func (u *Usecase) GetLastResult(ctx context.Context) (int64, error) {
	return u.repo.GetLastResult(ctx)
}
