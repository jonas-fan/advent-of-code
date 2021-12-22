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

func nearest(distance []int, visited []bool) int {
	min := 0
	where := 0

	for i := range distance {
		if visited[i] || distance[i] == 0 {
			continue
		}

		if min == 0 || distance[i] < min {
			min = distance[i]
			where = i
		}
	}

	return where
}

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

func solution(route [][]int, row int, col int) int {
	size := len(route)
	distance := make([]int, size*size)
	visited := make([]bool, size*size)

	distance[1], distance[size] = route[0][1], route[1][0]
	visited[0] = true

	for i := 0; i < len(distance)-1; i++ {
		if i%10000 == 0 {
			fmt.Printf("In progress ... (%d/%d)\n", i+1, len(distance))
		}

		from := nearest(distance, visited)
		fromRow, fromCol := from/size, from%size

		visited[from] = true

		for row := fromRow - 1; row < fromRow+2; row++ {
			for col := fromCol - 1; col < fromCol+2; col++ {
				to := row*size + col

				if row < 0 || row >= size || col < 0 || col >= size {
					// out of range
					continue
				} else if row == fromRow && col == fromCol {
					// same position
					continue
				} else if abs(row-fromRow)+abs(col-fromCol) > 1 {
					// cannot reach
					continue
				} else if visited[to] {
					// already checked
					continue
				}

				if distance[to] == 0 || distance[from]+route[row][col] < distance[to] {
					distance[to] = distance[from] + route[row][col]
				}
			}
		}
	}

	return distance[len(distance)-1]
}

func main() {
	route := [][]int{}

	for input := range read(os.Stdin) {
		line := []int{}

		for _, each := range input {
			line = append(line, int(each-'0'))
		}

		route = append(route, line)
	}

	route5x := make([][]int, len(route)*5)

	for row, line := range route {
		for i := 0; i < 5; i++ {
			route5x[row+i*len(route)] = make([]int, len(line)*5)
		}
	}

	for row, line := range route {
		for col, each := range line {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					val := ((each + i + j) + ((each + i + j) / 10)) % 10

					route5x[row+i*len(route)][col+j*len(line)] = val
				}
			}
		}
	}

	fmt.Println(solution(route, 0, 0))
	fmt.Println(solution(route5x, 0, 0))
}
