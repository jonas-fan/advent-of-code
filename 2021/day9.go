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

func height(heightmap [][]rune, row int, col int) int {
	if row < 0 || row >= len(heightmap) {
		return 9
	} else if col < 0 || col >= len(heightmap[0]) {
		return 9
	}

	return int(heightmap[row][col] - '0')
}

func dented(heightmap [][]rune, row int, col int) bool {
	current := height(heightmap, row, col)

	switch {
	case current >= height(heightmap, row-1, col):
		return false
	case current >= height(heightmap, row+1, col):
		return false
	case current >= height(heightmap, row, col-1):
		return false
	case current >= height(heightmap, row, col+1):
		return false
	}

	return true
}

func climb(heightmap [][]rune, row int, col int, visited [][]bool) int {
	if row < 0 || row >= len(heightmap) {
		return 0
	} else if col < 0 || col >= len(heightmap[0]) {
		return 0
	} else if visited[row][col] {
		return 0
	} else if heightmap[row][col] >= '9' {
		return 0
	}

	visited[row][col] = true

	return 1 +
		climb(heightmap, row-1, col, visited) +
		climb(heightmap, row+1, col, visited) +
		climb(heightmap, row, col-1, visited) +
		climb(heightmap, row, col+1, visited)
}

func solution1(heightmap [][]rune) int {
	risk := 0

	for row := range heightmap {
		for col := range heightmap[row] {
			if dented(heightmap, row, col) {
				risk += height(heightmap, row, col) + 1
			}
		}
	}

	return risk
}

func solution2(heightmap [][]rune) int {
	visited := make([][]bool, 0, len(heightmap))

	for _, each := range heightmap {
		visited = append(visited, make([]bool, len(each)))
	}

	basin := []int{}

	for row := range heightmap {
		for col := range heightmap[row] {
			if dented(heightmap, row, col) {
				basin = append(basin, climb(heightmap, row, col, visited))
			}
		}
	}

	sort.Slice(basin, func(lhs int, rhs int) bool {
		return basin[lhs] > basin[rhs]
	})

	out := 1

	for i := 0; i < len(basin) && i < 3; i++ {
		out *= basin[i]
	}

	return out
}

func main() {
	heightmap := [][]rune{}

	for input := range read(os.Stdin) {
		heightmap = append(heightmap, []rune(input))
	}

	fmt.Println(solution1(heightmap))
	fmt.Println(solution2(heightmap))
}
