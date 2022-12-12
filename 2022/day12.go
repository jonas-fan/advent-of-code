package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func min(lhs int, rhs int) int {
	if lhs < rhs {
		return lhs
	}

	return rhs
}

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}

	return lhs
}

func int2D(row int, col int) [][]int {
	out := make([][]int, row)

	for i := range out {
		out[i] = make([]int, col)
	}

	return out
}

func bool2D(row int, col int) [][]bool {
	out := make([][]bool, row)

	for i := range out {
		out[i] = make([]bool, col)
	}

	return out
}

type Point struct {
	row int
	col int
}

func nearest(distance [][]int, visited [][]bool) *Point {
	min := 0
	where := &Point{}

	for i := 0; i < len(distance); i++ {
		for j := 0; j < len(distance[i]); j++ {
			if visited[i][j] || distance[i][j] == 0 {
				continue
			}

			if min == 0 || distance[i][j] < min {
				min = distance[i][j]
				where.row = i
				where.col = j
			}
		}
	}

	return where
}

func sol1(start *Point, end *Point, route [][]int) int {
	height := len(route)
	width := len(route[0])
	distance := int2D(height, width)
	visited := bool2D(height, width)

	visited[start.row][start.col] = true

	// figure out nearby position
	for row := max(start.row-1, 0); row < min(start.row+2, height); row++ {
		for col := max(start.col-1, 0); col < min(start.col+2, width); col++ {
			if abs(row-start.row)+abs(col-start.col) > 1 {
				// cannot reach
				continue
			} else if route[row][col] > route[start.row][start.col]+1 {
				// too far
				continue
			}

			if row == start.row && col == start.col {
				distance[row][col] = 0
			} else {
				distance[row][col] = 1
			}
		}
	}

	// figure out the shorest path for each position
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			from := nearest(distance, visited)

			visited[from.row][from.col] = true

			// relax
			for row := max(from.row-1, 0); row < min(from.row+2, height); row++ {
				for col := max(from.col-1, 0); col < min(from.col+2, width); col++ {
					to := &Point{
						row: row,
						col: col,
					}

					if visited[to.row][to.col] {
						// already checked
						continue
					} else if to.row == from.row && to.col == from.col {
						// same position
						continue
					} else if abs(to.row-from.row)+abs(to.col-from.col) > 1 {
						// cannot reach
						continue
					} else if route[to.row][to.col] > route[from.row][from.col]+1 {
						// too far
						continue
					}

					if distance[to.row][to.col] == 0 || distance[from.row][from.col]+1 < distance[to.row][to.col] {
						distance[to.row][to.col] = distance[from.row][from.col] + 1
					}
				}
			}
		}
	}

	return distance[end.row][end.col]
}

func sol2(end *Point, route [][]int) int {
	out := math.MaxInt

	for row := range route {
		for col := range route {
			if route[row][col] > 0 {
				continue
			}

			start := &Point{
				row: row,
				col: col,
			}

			out = min(out, sol1(start, end, route))
		}
	}

	return out
}

func main() {
	lines := []string{}

	for input := range read(os.Stdin) {
		lines = append(lines, input)
	}

	start := &Point{}
	end := &Point{}
	route := int2D(len(lines), len(lines[0]))

	for row, line := range lines {
		for col, letter := range line {
			if letter == 'S' {
				letter = 'a'
				start.row = row
				start.col = col
			} else if letter == 'E' {
				letter = 'z'
				end.row = row
				end.col = col
			}

			route[row][col] = int(letter - 'a')
		}
	}

	fmt.Println("Ans1", sol1(start, end, route))
	fmt.Println("Ans2", sol2(end, route))
}
