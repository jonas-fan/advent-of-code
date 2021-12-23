package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
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

func solution1(points []int) int {
	ymin := min(points[2], points[3])
	vy := -ymin - 1

	return (1 + vy) * vy / 2
}

func solution2(points []int) int {
	xmin := min(points[0], points[1])
	xmax := max(points[0], points[1])
	ymin := min(points[2], points[3])
	ymax := max(points[2], points[3])
	reaches := 0

	for x := 1; x <= xmax; x++ {
		for y := ymin; y <= -ymin-1; y++ {
			xpos, ypos := x, y
			xdelta, ydelta := x, y

			for {
				if xpos >= xmin && xpos <= xmax && ypos >= ymin && ypos <= ymax {
					reaches++
					break
				}

				xdelta = max(xdelta-1, 0)
				ydelta--
				xpos += xdelta
				ypos += ydelta

				if ypos < ymin {
					break
				} else if xpos > xmax {
					break
				} else if xdelta == 0 && xpos < xmin {
					break
				}
			}
		}
	}

	return reaches
}

func main() {
	token := ""

	for input := range read(os.Stdin) {
		token = input
	}

	re := regexp.MustCompile(`target area: x=([^.]+)..([^.]+), y=([^.]+)..([^.]+)`)
	matches := re.FindStringSubmatch(token)

	fmt.Println(solution1(str2intSlice(matches[1:])))
	fmt.Println(solution2(str2intSlice(matches[1:])))
}
