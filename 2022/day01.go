package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read(reader io.Reader) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			out <- strings.TrimSpace(scanner.Text())
		}
	}()

	return out
}

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}

	return lhs
}

func sum(nums ...int) int {
	out := 0

	for _, num := range nums {
		out += num
	}

	return out
}

func str2int(str string) int {
	val, _ := strconv.Atoi(str)

	return val
}

func solution1(bucket [][]int) int {
	out := 0

	for _, nums := range bucket {
		out = max(out, sum(nums...))
	}

	return out
}

func solution2(bucket [][]int) int {
	calories := make([]int, len(bucket))

	for i, nums := range bucket {
		calories[i] = sum(nums...)
	}

	sort.Ints(calories)

	begin := max(0, len(calories)-3)

	return sum(calories[begin:]...)
}

func main() {
	nums := make([][]int, 1)
	index := 0

	for input := range read(os.Stdin) {
		if input == "" {
			nums = append(nums, make([]int, 0))
			index++
		} else {
			nums[index] = append(nums[index], str2int(input))
		}
	}

	fmt.Println(solution1(nums))
	fmt.Println(solution2(nums))
}
