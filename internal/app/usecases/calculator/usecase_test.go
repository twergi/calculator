package calculator

import (
	"github.com/twergi/calculator/tests/mocks"
	"go.uber.org/mock/gomock"
)

var (
	_any = gomock.Any()
)

func newMockedUsecase(m *mocks.Mocker) *Usecase {
	return New(m.MockRepository)
}
