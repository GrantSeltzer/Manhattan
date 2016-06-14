package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/docker/engine-api/types"
)

func parseArguments(needed bool, delimArgs []string, syscallStruct *types.Syscall) {
	syscallIndex, _ := strconv.ParseUint(delimArgs[1], 10, 0)
	syscallValue, _ := strconv.ParseUint(delimArgs[2], 10, 64)
	syscallValueTwo, _ := strconv.ParseUint(delimArgs[3], 10, 64)
	syscallOp := parseOperator(delimArgs[4])

	ArgStruct := types.Arg{
		Index:    uint(syscallIndex),
		Value:    syscallValue,
		ValueTwo: syscallValueTwo,
		Op:       syscallOp,
	}
	var ArgSlice []*types.Arg
	ArgSlice = append(ArgSlice, &ArgStruct)
	syscallStruct.Args = ArgSlice
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
		return types.OpNotEqual
	}
}
