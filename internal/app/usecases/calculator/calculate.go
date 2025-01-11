package calculator

import (
	"context"
	"errors"
	"math"

	"github.com/twergi/calculator/internal/model"
)

var errOverflow = errors.New("overflow")

func (u *Usecase) Calculate(ctx context.Context, a, b int64, operation model.OperationType) (result int64, err error) {

	switch operation {
	case model.OperationTypeSum:
		result, err = u.sum(a, b)
	case model.OperationTypeSub:
		result, err = u.sub(a, b)
	case model.OperationTypeMult:
		result, err = u.mult(a, b)
	case model.OperationTypeDiv:
		result, err = u.div(a, b)
	case model.OperationTypeMod:
		result = a % b
	default:
		err = errors.New("invalid operation type")
	}
	if err != nil {
		return 0, err
	}

	err = u.repo.SaveResult(ctx, result)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func (u *Usecase) sum(a, b int64) (int64, error) {
	if b > 0 {
		if a > math.MaxInt64-b {
			return 0, errOverflow
		}
	} else {
		if a < math.MinInt64-b {
			return 0, errOverflow
		}
	}
	return a + b, nil
}

func (u *Usecase) sub(a, b int64) (int64, error) {
	if b > 0 {
		if a < math.MinInt64+b {
			return 0, errOverflow
		}
	} else {
		if a > math.MaxInt64+b {
			return 0, errOverflow
		}
	}

	return a - b, nil
}

func (u *Usecase) mult(a, b int64) (int64, error) {
	result := a * b
	if (result < 0) == ((a < 0) != (b < 0)) {
		if result/b == a {
			return result, nil
		}
	}
	return 0, errOverflow
}

func (u *Usecase) div(a, b int64) (int64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by 0")
	}

	return a / b, nil
}
