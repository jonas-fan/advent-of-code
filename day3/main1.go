package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func solution(binaries [][]int) int {
	if len(binaries) == 0 {
		return 0
	}

	bits := make([]int, len(binaries[0]))

	for _, each := range binaries {
		for i, bit := range each {
			bits[i] += bit
		}
	}

	for i, each := range bits {
		if each < (len(binaries) >> 1) {
			bits[i] = 0
		} else {
			bits[i] = 1
		}
	}

	gamma := 0
	epsilon := 0

	for _, each := range bits {
		gamma = (gamma << 1) | each
		epsilon = (epsilon << 1) | (1 - each)
	}

	return gamma * epsilon
}

func main() {
	binaries := make([][]int, 0)

	for input := range read(os.Stdin) {
		binary := make([]int, len(input))

		for i, digit := range input {
			binary[i] = int(digit - '0')
		}

		binaries = append(binaries, binary)
	}

	fmt.Println(solution(binaries))
}
