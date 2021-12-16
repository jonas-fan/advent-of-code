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

func solution(nums []int) int {
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
		num, err := strconv.Atoi(input)

		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	fmt.Println(solution(nums))
}
