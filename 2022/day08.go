package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func read(reader io.Reader) chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			out <- scanner.Text()
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

func digit2D(lines []string) [][]int {
	out := make([][]int, 0)

	for _, line := range lines {
		row := []int{}

		for _, col := range line {
			row = append(row, int(col-'0'))
		}

		out = append(out, row)
	}

	return out
}

func visableUp(trees [][]int, row int, col int) (bool, int) {
	out := 0
	val := trees[row][col]

	for row = row - 1; row >= 0; row-- {
		out++

		if val <= trees[row][col] {
			return false, out
		}
	}

	return true, out
}

func visableDown(trees [][]int, row int, col int) (bool, int) {
	out := 0
	val := trees[row][col]

	for row = row + 1; row < len(trees); row++ {
		out++

		if val <= trees[row][col] {
			return false, out
		}
	}

	return true, out
}

func visableLeft(trees [][]int, row int, col int) (bool, int) {
	out := 0
	val := trees[row][col]

	for col = col - 1; col >= 0; col-- {
		out++

		if val <= trees[row][col] {
			return false, out
		}
	}

	return true, out
}

func visableRight(trees [][]int, row int, col int) (bool, int) {
	out := 0
	val := trees[row][col]

	for col = col + 1; col < len(trees[row]); col++ {
		out++

		if val <= trees[row][col] {
			return false, out
		}
	}

	return true, out
}

func sol1(trees [][]int) int {
	out := 0

	for row := 0; row < len(trees); row++ {
		for col := 0; col < len(trees[row]); col++ {
			var ok bool

			if ok, _ = visableUp(trees, row, col); ok {
				out++
			} else if ok, _ = visableDown(trees, row, col); ok {
				out++
			} else if ok, _ = visableLeft(trees, row, col); ok {
				out++
			} else if ok, _ = visableRight(trees, row, col); ok {
				out++
			}
		}
	}

	return out
}

func sol2(trees [][]int) int {
	out := 0

	for row := 0; row < len(trees); row++ {
		for col := 0; col < len(trees[row]); col++ {
			_, up := visableUp(trees, row, col)
			_, down := visableDown(trees, row, col)
			_, left := visableLeft(trees, row, col)
			_, right := visableRight(trees, row, col)

			out = max(out, up*down*left*right)
		}
	}

	return out
}

func main() {
	lines := []string{}

	for input := range read(os.Stdin) {
		lines = append(lines, input)
	}

	trees := digit2D(lines)

	fmt.Println("Ans1", sol1(trees))
	fmt.Println("Ans2", sol2(trees))
}
