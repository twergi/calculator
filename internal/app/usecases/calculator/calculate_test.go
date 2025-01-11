package calculator

import (
	"context"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/twergi/calculator/internal/model"
	"github.com/twergi/calculator/tests/mocks"
)

func TestUsecase_sum(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name    string
		setup   func(m *mocks.Mocker)
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:  "1+1=2",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: 1,
				b: 1,
			},
			want:    2,
			wantErr: false,
		},
		{
			name:  "(-1)+(-1)=-2",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: -1,
				b: -1,
			},
			want:    -2,
			wantErr: false,
		},
		{
			name:  "MaxInt64+(-MaxInt64)=0",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MaxInt64,
				b: -math.MaxInt64,
			},
			want:    0,
			wantErr: false,
		},
		{
			name:  "MinInt64+(-(MinInt64+1))=-1",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MinInt64,
				b: -(math.MinInt64 + 1),
			},
			want:    -1,
			wantErr: false,
		},
		{
			name:  "MinInt64+(-1)=overflow",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MinInt64,
				b: -1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name:  "-1+MinInt64=overflow",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: -1,
				b: math.MinInt64,
			},
			want:    0,
			wantErr: true,
		},
		{
			name:  "MaxInt64+1=overflow",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MaxInt64,
				b: 1,
			},
			want:    0,
			wantErr: true,
		},
		{
			name:  "1+MaxInt64=overflow",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: 1,
				b: math.MaxInt64,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := mocks.NewMocker(t)
			defer m.Finish()
			tt.setup(m)
			u := newMockedUsecase(m)

			got, err := u.sum(tt.args.a, tt.args.b)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUsecase_sub(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name    string
		setup   func(m *mocks.Mocker)
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:  "1-1=0",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: 1,
				b: 1,
			},
			want:    0,
			wantErr: false,
		},
		{
			name:  "(-1)-(-1)=0",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: -1,
				b: -1,
			},
			want:    0,
			wantErr: false,
		},
		{
			name:  "MaxInt64-(-MaxInt64)=overflow",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MaxInt64,
				b: -math.MaxInt64,
			},
			want:    0,
			wantErr: true,
		},
		{
			name:  "MinInt64-(-(MinInt64+1))=overflow",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MinInt64,
				b: -(math.MinInt64 + 1),
			},
			want:    0,
			wantErr: true,
		},
		{
			name:  "MinInt64-(-1)=MinInt64+1",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MinInt64,
				b: -1,
			},
			want:    math.MinInt64 + 1,
			wantErr: false,
		},
		{
			name:  "-1-MinInt64=MaxInt64",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: -1,
				b: math.MinInt64,
			},
			want:    math.MaxInt64,
			wantErr: false,
		},
		{
			name:  "MaxInt64-1=MaxInt64-1",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MaxInt64,
				b: 1,
			},
			want:    math.MaxInt64 - 1,
			wantErr: false,
		},
		{
			name:  "1-MaxInt64=MinInt64+2",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: 1,
				b: math.MaxInt64,
			},
			want:    math.MinInt64 + 2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := mocks.NewMocker(t)
			defer m.Finish()
			tt.setup(m)
			u := newMockedUsecase(m)

			got, err := u.sub(tt.args.a, tt.args.b)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUsecase_mult(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name    string
		setup   func(m *mocks.Mocker)
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:  "1*1=1",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: 1,
				b: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name:  "2*2=4",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: 2,
				b: 2,
			},
			want:    4,
			wantErr: false,
		},
		{
			name:  "(-2)*2=-4",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: -2,
				b: 2,
			},
			want:    -4,
			wantErr: false,
		},
		{
			name:  "(-2)*(-2)=4",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: -2,
				b: -2,
			},
			want:    4,
			wantErr: false,
		},
		{
			name:  "MaxInt64*2=overflow",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MaxInt64,
				b: 2,
			},
			want:    0,
			wantErr: true,
		},
		{
			name:  "MaxInt64/2*2=MaxInt64-1",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MaxInt64 / 2,
				b: 2,
			},
			want:    math.MaxInt64 - 1,
			wantErr: false,
		},
		{
			name:  "MinInt64/2*2=MinInt64",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MinInt64 / 2,
				b: 2,
			},
			want:    math.MinInt64,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := mocks.NewMocker(t)
			defer m.Finish()
			tt.setup(m)
			u := newMockedUsecase(m)

			got, err := u.mult(tt.args.a, tt.args.b)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUsecase_div(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name    string
		setup   func(m *mocks.Mocker)
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:  "1/1=1",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: 1,
				b: 1,
			},
			want:    1,
			wantErr: false,
		},
		{
			name:  "1/2=0",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: 1,
				b: 2,
			},
			want:    0,
			wantErr: false,
		},
		{
			name:  "0/2=0",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: 0,
				b: 2,
			},
			want:    0,
			wantErr: false,
		},
		{
			name:  "1/MaxInt64=0",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: 1,
				b: math.MaxInt64,
			},
			want:    0,
			wantErr: false,
		},
		{
			name:  "MaxInt64/MaxInt64=1",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MaxInt64,
				b: math.MaxInt64,
			},
			want:    1,
			wantErr: false,
		},
		{
			name:  "MaxInt64/MinInt64=0",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MaxInt64,
				b: math.MinInt64,
			},
			want:    0,
			wantErr: false,
		},
		{
			name:  "MinInt64/MaxInt64=-1",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MinInt64,
				b: math.MaxInt64,
			},
			want:    -1,
			wantErr: false,
		},
		{
			name:  "MinInt64/0=error",
			setup: func(m *mocks.Mocker) {},
			args: args{
				a: math.MinInt64,
				b: 0,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := mocks.NewMocker(t)
			defer m.Finish()
			tt.setup(m)
			u := newMockedUsecase(m)

			got, err := u.div(tt.args.a, tt.args.b)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUsecase_Calculate(t *testing.T) {
	type args struct {
		ctx       context.Context
		a         int64
		b         int64
		operation model.OperationType
	}
	tests := []struct {
		name       string
		setup      func(m *mocks.Mocker)
		args       args
		wantResult int64
		wantErr    bool
	}{
		{
			name: "16+16",
			setup: func(m *mocks.Mocker) {
				m.MockRepository.EXPECT().
					SaveResult(_any, int64(32)).
					Return(nil)
			},
			args: args{
				ctx:       context.Background(),
				a:         16,
				b:         16,
				operation: model.OperationTypeSum,
			},
			wantResult: 32,
			wantErr:    false,
		},
		{
			name: "16-16",
			setup: func(m *mocks.Mocker) {
				m.MockRepository.EXPECT().
					SaveResult(_any, int64(0)).
					Return(nil)
			},
			args: args{
				ctx:       context.Background(),
				a:         16,
				b:         16,
				operation: model.OperationTypeSub,
			},
			wantResult: 0,
			wantErr:    false,
		},
		{
			name: "16*16",
			setup: func(m *mocks.Mocker) {
				m.MockRepository.EXPECT().
					SaveResult(_any, int64(256)).
					Return(nil)
			},
			args: args{
				ctx:       context.Background(),
				a:         16,
				b:         16,
				operation: model.OperationTypeMult,
			},
			wantResult: 256,
			wantErr:    false,
		},
		{
			name: "16/16",
			setup: func(m *mocks.Mocker) {
				m.MockRepository.EXPECT().
					SaveResult(_any, int64(1)).
					Return(nil)
			},
			args: args{
				ctx:       context.Background(),
				a:         16,
				b:         16,
				operation: model.OperationTypeDiv,
			},
			wantResult: 1,
			wantErr:    false,
		},
		{
			name: "16%16",
			setup: func(m *mocks.Mocker) {
				m.MockRepository.EXPECT().
					SaveResult(_any, int64(0)).
					Return(nil)
			},
			args: args{
				ctx:       context.Background(),
				a:         16,
				b:         16,
				operation: model.OperationTypeMod,
			},
			wantResult: 0,
			wantErr:    false,
		},
		{
			name: "15%16",
			setup: func(m *mocks.Mocker) {
				m.MockRepository.EXPECT().
					SaveResult(_any, int64(15)).
					Return(nil)
			},
			args: args{
				ctx:       context.Background(),
				a:         15,
				b:         16,
				operation: model.OperationTypeMod,
			},
			wantResult: 15,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := mocks.NewMocker(t)
			defer m.Finish()
			tt.setup(m)
			u := newMockedUsecase(m)

			got, err := u.Calculate(tt.args.ctx, tt.args.a, tt.args.b, tt.args.operation)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantResult, got)
		})
	}
}
