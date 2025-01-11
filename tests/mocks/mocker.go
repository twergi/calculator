package mocks

import (
	"testing"

	gomock "go.uber.org/mock/gomock"
)

type Mocker struct {
	ctrl           *gomock.Controller
	MockRepository *MockRepository
}

func NewMocker(t *testing.T) *Mocker {
	ctrl := gomock.NewController(t)

	return &Mocker{
		ctrl:           ctrl,
		MockRepository: NewMockRepository(ctrl),
	}
}

func (m *Mocker) Finish() {
	m.ctrl.Finish()
}
