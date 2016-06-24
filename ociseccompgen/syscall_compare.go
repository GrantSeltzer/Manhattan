package ociseccompgen

import (
	"reflect"

	types "github.com/opencontainers/runtime-spec/specs-go"
)

func compareSyscalls(config1, config2 *types.Syscall) bool {
	totalSimilarities := 0

	if config1.Name == config2.Name {
		totalSimilarities++
	}

	if config1.Action == config2.Action {
		totalSimilarities++
	}

	if reflect.DeepEqual(config1.Args, config2.Args) {
		totalSimilarities++
	}

	if totalSimilarities == 3 {
		return true
	}
	return false
}
