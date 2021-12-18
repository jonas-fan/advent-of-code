package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Move struct {
	from Point
	to   Point
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

func min(lhs int, rhs int) int {
	if lhs < rhs {
		return lhs
	}

	return rhs
}

func solution(moves []*Move) int {
	visited := make(map[string]int)

	for _, move := range moves {
		from, to := &move.from, &move.to
		stepX, stepY := 0, 0

		if from.x < to.x {
			stepX = 1
		} else if from.x > to.x {
			stepX = -1
		}

		if from.y < to.y {
			stepY = 1
		} else if from.y > to.y {
			stepY = -1
		}

		if stepX != 0 && stepY != 0 {
			continue
		}

		current := from

		for current.x != to.x || current.y != to.y {
			key := fmt.Sprintf("%s,%s", current.x, current.y)

			visited[key]++
			current.x += stepX
			current.y += stepY
		}

		key := fmt.Sprintf("%s,%s", to.x, to.y)

		visited[key]++
	}

	overlapped := 0

	for _, count := range visited {
		if count > 1 {
			overlapped++
		}
	}

	return overlapped
}

func main() {
	moves := []*Move{}

	for input := range read(os.Stdin) {
		points := strings.Split(input, " -> ")
		from := strings.Split(points[0], ",")
		to := strings.Split(points[1], ",")
		move := &Move{
			from: Point{x: str2int(from[0]), y: str2int(from[1])},
			to:   Point{x: str2int(to[0]), y: str2int(to[1])},
		}

		moves = append(moves, move)
	}

	fmt.Println(solution(moves))
}
