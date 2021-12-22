package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Position struct {
	row int
	col int
}

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

func spread(octopuses [][]int, row int, col int) int {
	if row < 0 || row >= len(octopuses) {
		return 0
	} else if col < 0 || col >= len(octopuses[0]) {
		return 0
	} else if octopuses[row][col] == 0 {
		return 0
	}

	octopuses[row][col] = (octopuses[row][col] + 1) % 10

	if octopuses[row][col] != 0 {
		return 0
	}

	return 1 +
		spread(octopuses, row-1, col-1) +
		spread(octopuses, row-1, col) +
		spread(octopuses, row-1, col+1) +
		spread(octopuses, row, col-1) +
		spread(octopuses, row, col+1) +
		spread(octopuses, row+1, col-1) +
		spread(octopuses, row+1, col) +
		spread(octopuses, row+1, col+1)
}

func solution1(octopuses [][]int, steps int) int {
	flashes := 0

	for ; steps > 0; steps-- {
		ready := []*Position{}

		for row, line := range octopuses {
			for col, each := range line {
				octopuses[row][col] = (each + 1) % 10

				if octopuses[row][col] == 0 {
					ready = append(ready, &Position{row: row, col: col})
				}
			}
		}

		for _, pos := range ready {
			octopuses[pos.row][pos.col] = 9

			flashes += spread(octopuses, pos.row, pos.col)
		}
	}

	return flashes
}

func solution2(octopuses [][]int) int {
	steps := 0
	flashes := 0
	count := len(octopuses) * len(octopuses[0])

	for steps = 0; flashes != count; steps++ {
		ready := []*Position{}

		for row, line := range octopuses {
			for col, each := range line {
				octopuses[row][col] = (each + 1) % 10

				if octopuses[row][col] == 0 {
					ready = append(ready, &Position{row: row, col: col})
				}
			}
		}

		flashes = 0

		for _, pos := range ready {
			octopuses[pos.row][pos.col] = 9

			flashes += spread(octopuses, pos.row, pos.col)
		}
	}

	return steps
}

func main() {
	octopuses := [][]int{}

	for input := range read(os.Stdin) {
		line := make([]int, 0, len(input))

		for _, each := range input {
			line = append(line, int(each-'0'))
		}

		octopuses = append(octopuses, line)
	}

	fmt.Println(solution1(octopuses, 100))
	fmt.Println(solution2(octopuses) + 100)
}
