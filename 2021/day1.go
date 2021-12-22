package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func str2int(str string) int {
	val, _ := strconv.Atoi(str)

	return val
}

func solution1(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	increased := 0
	last := 0

	for _, each := range nums {
		if each > last {
			increased++
		}

		last = each
	}

	return increased - 1
}

func solution2(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	increased := 0
	last := 0
	sum := nums[0] + nums[1]

	for i := 2; i < len(nums); i++ {
		sum += nums[i]

		if sum > last {
			increased++
		}

		last = sum
		sum -= nums[i-2]
	}

	return increased - 1
}

func main() {
	nums := make([]int, 0)

	for input := range read(os.Stdin) {
		nums = append(nums, str2int(input))
	}

	fmt.Println(solution1(nums))
	fmt.Println(solution2(nums))
}
