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

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

func solution1(position []int) int {
	pos := make(map[int]int)

	for _, each := range position {
		pos[each]++
	}

	min := -1

	for i := range pos {
		fuel := 0

		for j, count := range pos {
			if i != j {
				fuel += abs(i-j) * count
			}
		}

		if min < 0 || fuel < min {
			min = fuel
		}
	}

	return min
}

func solution2(position []int) int {
	pos := make(map[int]int)

	for _, each := range position {
		pos[each]++
	}

	min := -1

	for i := range pos {
		fuel := 0

		for j, count := range pos {
			if i != j {
				dist := abs(i - j)

				fuel += count * (((1 + dist) * dist) >> 1)
			}
		}

		if min < 0 || fuel < min {
			min = fuel
		}
	}

	return min
}

func main() {
	position := []int{}

	for input := range read(os.Stdin) {
		for _, each := range strings.Split(input, ",") {
			position = append(position, str2int(each))
		}
	}

	fmt.Println(solution1(position))
	fmt.Println(solution2(position))
}
