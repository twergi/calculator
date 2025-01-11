package model

type OperationType int

const (
	OperationTypeInvalid OperationType = iota
	OperationTypeSum
	OperationTypeSub
	OperationTypeDiv
	OperationTypeMult
	OperationTypeMod
)

func OperationFromString(op string) OperationType {
	switch op {
	case "+":
		return OperationTypeSum
	case "-":
		return OperationTypeSub
	case "*":
		return OperationTypeMult
	case "/":
		return OperationTypeDiv
	case "%":
		return OperationTypeMod
	}

	return OperationTypeInvalid
}
