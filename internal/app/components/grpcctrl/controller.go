package grpcctrl

import (
	"context"

	"github.com/twergi/calculator/internal/app/usecases/calculator"
	"github.com/twergi/calculator/internal/model"
	desc "github.com/twergi/calculator/internal/proto/gen/go/service"
)

type Implementation struct {
	desc.UnimplementedCalculatorServer

	calculatorUsecase *calculator.Usecase
}

func New(calcUsecase *calculator.Usecase) *Implementation {
	return &Implementation{
		calculatorUsecase: calcUsecase,
	}
}

func (i *Implementation) Calculate(ctx context.Context, req *desc.CalculateRequest) (*desc.CalculateResponse, error) {
	res, err := i.calculatorUsecase.Calculate(ctx, req.A, req.B, mapOpFromPB(req.Operation))
	if err != nil {
		return nil, err
	}

	return &desc.CalculateResponse{
		Result: res,
	}, nil
}

func mapOpFromPB(op desc.CalculateOperationEnum) model.OperationType {
	switch op {
	case desc.CalculateOperationEnum_CALCULATE_OPERATION_SUM:
		return model.OperationTypeSum
	case desc.CalculateOperationEnum_CALCULATE_OPERATION_SUB:
		return model.OperationTypeSub
	case desc.CalculateOperationEnum_CALCULATE_OPERATION_MULT:
		return model.OperationTypeMult
	case desc.CalculateOperationEnum_CALCULATE_OPERATION_DIV:
		return model.OperationTypeDiv
	case desc.CalculateOperationEnum_CALCULATE_OPERATION_MOD:
		return model.OperationTypeMod
	}

	return model.OperationTypeInvalid
}

func (i *Implementation) GetPrevious(ctx context.Context, req *desc.GetPreviousRequest) (*desc.GetPreviousResponse, error) {
	res, err := i.calculatorUsecase.GetLastResult(ctx)
	if err != nil {
		return nil, err
	}
	return &desc.GetPreviousResponse{
		Result: res,
	}, nil
}
