package funcs

import (
	"math"
)

func Abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func Min(nums ...int) int {
	out := math.MaxInt32

	for _, num := range nums {
		if num < out {
			out = num
		}
	}

	return out
}

func Max(nums ...int) int {
	out := 0

	for _, num := range nums {
		if num > out {
			out = num
		}
	}

	return out
}


