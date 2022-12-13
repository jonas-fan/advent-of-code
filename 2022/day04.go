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

func str2intArray(strs []string) []int {
	out := make([]int, 0, len(strs))

	for _, str := range strs {
		out = append(out, str2int(str))
	}

	return out
}

func solution1(pairs [][]string) int {
	out := 0

	for _, pair := range pairs {
		lsection := str2intArray(strings.Split(pair[0], "-"))
		rsection := str2intArray(strings.Split(pair[1], "-"))

		if lsection[0] <= rsection[0] && rsection[1] <= lsection[1] {
			out++
		} else if rsection[0] <= lsection[0] && lsection[1] <= rsection[1] {
			out++
		}
	}

	return out
}

func solution2(pairs [][]string) int {
	out := 0

	for _, pair := range pairs {
		lsection := str2intArray(strings.Split(pair[0], "-"))
		rsection := str2intArray(strings.Split(pair[1], "-"))

		if lsection[1] < rsection[0] || lsection[0] > rsection[1] {
			continue
		} else {
			out++
		}
	}

	return out
}

func main() {
	inputs := make([][]string, 0)

	for input := range read(os.Stdin) {
		inputs = append(inputs, strings.Split(input, ","))
	}

	fmt.Println(solution1(inputs))
	fmt.Println(solution2(inputs))
}
