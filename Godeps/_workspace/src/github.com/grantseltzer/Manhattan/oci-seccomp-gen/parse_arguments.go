package ociseccompgen

import (
	"errors"
	"fmt"
	"strconv"

	types "github.com/opencontainers/runtime-spec/specs-go"
)

// parseArguments takes a list of arguments (delimArgs)  and a pointer to a
// corresponding syscall struct. It parses and fills out the argument information
func parseArguments(delimArgs []string) (*[]types.Arg, error) {
	nilArgSlice := new([]types.Arg)

	if len(delimArgs) == 1 {
		return nilArgSlice, nil
	}

	if len(delimArgs) == 5 {
		syscallIndex, err := strconv.ParseUint(delimArgs[1], 10, 0)
		if err != nil {
			return nilArgSlice, err
		}

		syscallValue, err := strconv.ParseUint(delimArgs[2], 10, 64)
		if err != nil {
			return nilArgSlice, err
		}

		syscallValueTwo, err := strconv.ParseUint(delimArgs[3], 10, 64)
		if err != nil {
			return nilArgSlice, err
		}

		syscallOp, err := parseOperator(delimArgs[4])
		if err != nil {
			return nilArgSlice, err
		}

		argStruct := types.Arg{
			Index:    uint(syscallIndex),
			Value:    syscallValue,
			ValueTwo: syscallValueTwo,
			Op:       syscallOp,
		}

		argSlice := new([]types.Arg)
		*argSlice = append(*argSlice, argStruct)
		return argSlice, nil
	}

	return nilArgSlice, errors.New("Incorrect number of arguments passed with syscall")
}

func parseOperator(operator string) (types.Operator, error) {

	operators := map[string]types.Operator{
		"NE": types.OpNotEqual,
		"LT": types.OpLessThan,
		"LE": types.OpLessEqual,
		"EQ": types.OpEqualTo,
		"GE": types.OpGreaterEqual,
		"GT": types.OpGreaterThan,
		"ME": types.OpMaskedEqual,
	}
	for k, v := range operators {
		if operator == k {
			return v, nil
		}
	}
	return types.OpNotEqual, fmt.Errorf("Unrecognized operator: %s", operator)
}
