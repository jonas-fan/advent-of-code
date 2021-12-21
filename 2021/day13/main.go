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

func str2intSlice(strs []string) []int {
	out := make([]int, 0, len(strs))

	for _, each := range strs {
		out = append(out, str2int(each))
	}

	return out
}

func solution1(points [][]int, folds [][]int) int {
	paper := make(map[Point]bool)

	for _, point := range points {
		p := Point{
			x: point[0],
			y: point[1],
		}

		paper[p] = true
	}

	for _, fold := range folds {
		clean := []Point{}

		for point := range paper {
			if fold[0] > 0 && point.x > fold[0] {
				clean = append(clean, point)
				point.x = (fold[0] << 1) - point.x
			} else if fold[1] > 0 && point.y > fold[1] {
				clean = append(clean, point)
				point.y = (fold[1] << 1) - point.y
			}

			paper[point] = true
		}

		for _, point := range clean {
			delete(paper, point)
		}

		break
	}

	return len(paper)
}

func solution2(points [][]int, folds [][]int) int {
	paper := make(map[Point]bool)

	for _, point := range points {
		p := Point{
			x: point[0],
			y: point[1],
		}

		paper[p] = true
	}

	row := 0
	col := 0

	for _, fold := range folds {
		clean := []Point{}

		for point := range paper {
			if fold[0] > 0 && point.x > fold[0] {
				clean = append(clean, point)
				point.x = (fold[0] << 1) - point.x
				col = fold[0]
			} else if fold[1] > 0 && point.y > fold[1] {
				clean = append(clean, point)
				point.y = (fold[1] << 1) - point.y
				row = fold[1]
			}

			paper[point] = true
		}

		for _, point := range clean {
			delete(paper, point)
		}
	}

	bits := make([][]int, row)

	for i := range bits {
		bits[i] = make([]int, col)
	}

	for point := range paper {
		bits[point.y][point.x] = 1
	}

	for _, line := range bits {
		for _, each := range line {
			if each == 0 {
				fmt.Printf(". ")
			} else {
				fmt.Printf("# ")
			}
		}
		fmt.Println()
	}

	return 0
}

func main() {
	points := [][]int{}
	folds := [][]int{}

	for input := range read(os.Stdin) {
		tokensByComma := strings.Split(input, ",")
		tokensBySpace := strings.Split(input, " ")

		if len(tokensByComma) > 1 {
			points = append(points, str2intSlice(tokensByComma))
		} else if len(tokensBySpace) > 1 {
			tokens := strings.Split(tokensBySpace[len(tokensBySpace)-1], "=")

			switch tokens[0] {
			case "x":
				folds = append(folds, []int{str2int(tokens[1]), 0})
			case "y":
				folds = append(folds, []int{0, str2int(tokens[1])})
			}
		}
	}

	fmt.Println(solution1(points, folds))
	fmt.Println(solution2(points, folds))
}
