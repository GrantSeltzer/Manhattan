package parse

import (
	"fmt"
	"strconv"

	"github.com/docker/engine-api/types"
)

// Arguments takes a list of arguments (delimArgs)  and a pointer to a
// corresponding syscall struct. It parses and fills out the argument information
func Arguments(delimArgs []string, syscallStruct *types.Syscall) error {
	syscallIndex, err := strconv.ParseUint(delimArgs[1], 10, 0)
	if err != nil {
		return err
	}

	syscallValue, err := strconv.ParseUint(delimArgs[2], 10, 64)
	if err != nil {
		return err
	}

	syscallValueTwo, err := strconv.ParseUint(delimArgs[3], 10, 64)
	if err != nil {
		return err
	}

	syscallOp, err := parseOperator(delimArgs[4])
	if err != nil {
		return err
	}

	argStruct := types.Arg{
		Index:    uint(syscallIndex),
		Value:    syscallValue,
		ValueTwo: syscallValueTwo,
		Op:       syscallOp,
	}
	var argSlice []*types.Arg
	argSlice = append(argSlice, &argStruct)
	syscallStruct.Args = argSlice
	return nil
}

func parseOperator(operator string) (types.Operator, error) {
	switch operator {
	case "NE":
		return types.OpNotEqual, nil
	case "LT":
		return types.OpLessThan, nil
	case "LE":
		return types.OpLessEqual, nil
	case "EQ":
		return types.OpEqualTo, nil
	case "GE":
		return types.OpGreaterEqual, nil
	case "GT":
		return types.OpGreaterThan, nil
	case "ME":
		return types.OpMaskedEqual, nil
	default:
		return types.OpNotEqual, fmt.Errorf("Unrecognized operator: %s", operator)
	}
}
