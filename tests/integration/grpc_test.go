package integration

import (
	"context"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	desc "github.com/twergi/calculator/internal/proto/gen/go/service"
)

func TestCalculate(t *testing.T) {
	grpc := newGRPC()
	ctx := context.Background()

	t.Run("sum", func(t *testing.T) {
		defer truncate()

		t.Run("64+8=72", func(t *testing.T) {
			resp, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         64,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_SUM,
				B:         8,
			})
			require.NoError(t, err)

			assert.Equal(t, int64(72), resp.Result)
		})

		t.Run("0+0=0", func(t *testing.T) {
			resp, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         0,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_SUM,
				B:         0,
			})
			require.NoError(t, err)

			assert.Equal(t, int64(0), resp.Result)
		})

		t.Run("-1000000+5=-9999995", func(t *testing.T) {
			resp, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         -1000000,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_SUM,
				B:         5,
			})
			require.NoError(t, err)

			assert.Equal(t, int64(-999995), resp.Result)
		})

		t.Run("overflow error", func(t *testing.T) {
			_, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         math.MaxInt64,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_SUM,
				B:         1,
			})

			assert.Error(t, err)
		})
	})

	t.Run("sub", func(t *testing.T) {
		defer truncate()

		t.Run("64-8=56", func(t *testing.T) {
			resp, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         64,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_SUB,
				B:         8,
			})
			require.NoError(t, err)

			assert.Equal(t, int64(56), resp.Result)
		})

		t.Run("overflow error", func(t *testing.T) {
			_, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         math.MinInt64,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_SUB,
				B:         1,
			})

			assert.Error(t, err)
		})
	})

	t.Run("mult", func(t *testing.T) {
		defer truncate()

		t.Run("8*8=64", func(t *testing.T) {
			resp, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         8,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_MULT,
				B:         8,
			})
			require.NoError(t, err)

			assert.Equal(t, int64(64), resp.Result)
		})

		t.Run("overflow error", func(t *testing.T) {
			_, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         math.MaxInt64 / 2,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_MULT,
				B:         3,
			})

			assert.Error(t, err)
		})
	})

	t.Run("div", func(t *testing.T) {
		defer truncate()

		t.Run("64/8=8", func(t *testing.T) {
			resp, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         64,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_DIV,
				B:         8,
			})
			require.NoError(t, err)

			assert.Equal(t, int64(8), resp.Result)
		})
	})

	t.Run("mod", func(t *testing.T) {
		defer truncate()

		t.Run("64%8=0", func(t *testing.T) {
			resp, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         64,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_MOD,
				B:         8,
			})
			require.NoError(t, err)

			assert.Equal(t, int64(0), resp.Result)
		})
	})
}

func TestGetPrevious(t *testing.T) {
	grpc := newGRPC()
	ctx := context.Background()

	t.Run("no previous on init", func(t *testing.T) {
		_, err := grpc.GetPrevious(ctx, &desc.GetPreviousRequest{})
		assert.Error(t, err)
	})

	t.Run("sum cycle", func(t *testing.T) {
		defer truncate()

		v := int64(1)

		for range 10 {
			respC, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         v,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_SUM,
				B:         v,
			})
			require.NoError(t, err)

			assert.Equal(t, v+v, respC.GetResult())

			respP, err := grpc.GetPrevious(ctx, &desc.GetPreviousRequest{})
			require.NoError(t, err)

			assert.Equal(t, respC.GetResult(), respP.GetResult())

			v = respC.Result
		}
	})

	t.Run("sub cycle", func(t *testing.T) {
		defer truncate()

		v1 := int64(1000)
		v2 := int64(10)

		for range 10 {
			respC, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         v1,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_SUB,
				B:         v2,
			})
			require.NoError(t, err)

			assert.Equal(t, v1-v2, respC.GetResult())

			respP, err := grpc.GetPrevious(ctx, &desc.GetPreviousRequest{})
			require.NoError(t, err)

			assert.Equal(t, respC.GetResult(), respP.GetResult())

			v1 = respC.Result
		}
	})

	t.Run("mult cycle", func(t *testing.T) {
		defer truncate()

		v1 := int64(10)
		v2 := int64(10)

		for range 10 {
			respC, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         v1,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_MULT,
				B:         v2,
			})
			require.NoError(t, err)

			assert.Equal(t, v1*v2, respC.GetResult())

			respP, err := grpc.GetPrevious(ctx, &desc.GetPreviousRequest{})
			require.NoError(t, err)

			assert.Equal(t, respC.GetResult(), respP.GetResult())

			v1 = respC.Result
		}
	})

	t.Run("div cycle", func(t *testing.T) {
		defer truncate()

		v1 := int64(1000000000)
		v2 := int64(10)

		for range 10 {
			respC, err := grpc.Calculate(ctx, &desc.CalculateRequest{
				A:         v1,
				Operation: desc.CalculateOperationEnum_CALCULATE_OPERATION_DIV,
				B:         v2,
			})
			require.NoError(t, err)

			assert.Equal(t, v1/v2, respC.GetResult())

			respP, err := grpc.GetPrevious(ctx, &desc.GetPreviousRequest{})
			require.NoError(t, err)

			assert.Equal(t, respC.GetResult(), respP.GetResult())

			v1 = respC.Result
		}
	})

}
