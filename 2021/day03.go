package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func solution1(binaries [][]int) int {
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

func b2i(bits []int) int {
	out := 0

	for _, each := range bits {
		out = (out << 1) | each
	}

	return out
}

func dfs(binaries [][]int, pos int, most bool) int {
	if len(binaries) == 0 {
		return 0
	} else if len(binaries) == 1 {
		return b2i(binaries[0])
	}

	sort.Slice(binaries, func(lhs int, rhs int) bool {
		if binaries[lhs][pos] > 0 && binaries[rhs][pos] == 0 {
			return true
		} else if binaries[lhs][pos] == 0 && binaries[rhs][pos] > 0 {
			return false
		}

		return binaries[lhs][pos] > 0 && binaries[rhs][pos] > 0
	})

	count := 1

	for binaries[count][pos] == binaries[count-1][pos] {
		count++
	}

	if most {
		if (count << 1) < len(binaries) {
			binaries = binaries[count:]
		} else {
			binaries = binaries[:count]
		}
	} else {
		if (count << 1) < len(binaries) {
			binaries = binaries[:count]
		} else {
			binaries = binaries[count:]
		}
	}

	return dfs(binaries, pos+1, most)
}

func solution2(binaries [][]int) int {
	if len(binaries) == 0 {
		return 0
	}

	return dfs(binaries, 0, true) * dfs(binaries, 0, false)
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

	fmt.Println(solution1(binaries))
	fmt.Println(solution2(binaries))
}
