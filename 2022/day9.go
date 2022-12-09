package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func atoi(str string) int {
	out, _ := strconv.Atoi(str)

	return out
}

type Direction struct {
	direct rune
	steps  int
}

type Point struct {
	x int
	y int
}

func sol(directs []*Direction, num int) int {
	knots := make([]Point, num)

	visited := map[Point]int{}
	visited[knots[0]]++

	for _, direct := range directs {
		for i := 0; i < direct.steps; i++ {
			// move the head
			switch direct.direct {
			case 'U':
				knots[0].y++
			case 'D':
				knots[0].y--
			case 'L':
				knots[0].x--
			case 'R':
				knots[0].x++
			}

			for j := 1; j < len(knots); j++ {
				head := knots[j-1]
				curr := knots[j]

				if head == curr {
					// cover
				} else if abs(head.x-curr.x) < 2 && abs(head.y-curr.y) < 2 {
					// touching
				} else if head.x == curr.x {
					// aligned
					if head.y > curr.y {
						curr.y++
					} else {
						curr.y--
					}
				} else if head.y == curr.y {
					// aligned
					if head.x > curr.x {
						curr.x++
					} else {
						curr.x--
					}
				} else {
					// not aligned
					if head.x > curr.x {
						curr.x++
					} else {
						curr.x--
					}

					if head.y > curr.y {
						curr.y++
					} else {
						curr.y--
					}
				}

				knots[j] = curr
			}

			// mark the tail
			visited[knots[len(knots)-1]]++
		}
	}

	return len(visited)
}

func main() {
	directions := []*Direction{}

	for input := range read(os.Stdin) {
		tokens := strings.Split(input, " ")
		direct := &Direction{
			direct: rune(tokens[0][0]),
			steps:  atoi(tokens[1]),
		}

		directions = append(directions, direct)
	}

	fmt.Println("Ans1", sol(directions, 2))
	fmt.Println("Ans2", sol(directions, 10))
}
