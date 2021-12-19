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
