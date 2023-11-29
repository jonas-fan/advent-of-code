package funcs

import (
	"strconv"
)

func Atoi(str string) int {
	out, _ := strconv.Atoi(str)

	return out
}

func Atois(strs []string) []int {
	out := make([]int, 0, len(strs))

	for _, each := range strs {
		out = append(out, Atoi(each))
	}

	return out
}

