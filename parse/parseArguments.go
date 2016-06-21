package parse

import (
	"fmt"
	"os"
	"strconv"

	"github.com/docker/engine-api/types"
)

// Arguments takes a list of arguments (delimArgs)  and a pointer to a
// corresponding syscall struct. It parses and fills out the argument information
func Arguments(delimArgs []string, syscallStruct *types.Syscall) {
	syscallIndex, _ := strconv.ParseUint(delimArgs[1], 10, 0)
	syscallValue, _ := strconv.ParseUint(delimArgs[2], 10, 64)
	syscallValueTwo, _ := strconv.ParseUint(delimArgs[3], 10, 64)
	syscallOp := parseOperator(delimArgs[4])

	argStruct := types.Arg{
		Index:    uint(syscallIndex),
		Value:    syscallValue,
		ValueTwo: syscallValueTwo,
		Op:       syscallOp,
	}
	argSlice := []*types.Arg{}
	argSlice = append(argSlice, &argStruct)
	syscallStruct.Args = argSlice
}

func parseOperator(operator string) types.Operator {
	switch operator {
	case "NE":
		return types.OpNotEqual
	case "LT":
		return types.OpLessThan
	case "LE":
		return types.OpLessEqual
	case "EQ":
		return types.OpEqualTo
	case "GE":
		return types.OpGreaterEqual
	case "GT":
		return types.OpGreaterThan
	case "ME":
		return types.OpMaskedEqual
	default:
		fmt.Println("Unrecognized operator", operator)
		os.Exit(-3)
	}
}
