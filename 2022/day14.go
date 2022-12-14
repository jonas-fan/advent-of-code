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

func atoi(str string) int {
	out, _ := strconv.Atoi(str)

	return out
}

func atois(strs []string) []int {
	out := make([]int, 0, len(strs))

	for _, each := range strs {
		out = append(out, atoi(each))
	}

	return out
}

type Coordinate struct {
	x int
	y int
}

func fall(sand *Coordinate, cave map[Coordinate]int, height int, infinite bool) *Coordinate {
	if infinite {
		if sand.y > height {
			return nil
		}
	} else if sand.y+1 == height {
		return sand
	}

	nexts := []*Coordinate{
		&Coordinate{sand.x, sand.y + 1},
		&Coordinate{sand.x - 1, sand.y + 1},
		&Coordinate{sand.x + 1, sand.y + 1},
	}

	for _, next := range nexts {
		if _, have := cave[*next]; !have {
			return fall(next, cave, height, infinite)
		}
	}

	return sand
}

func sol1(cave map[Coordinate]int, height int) int {
	sand := Coordinate{500, 0}

	for count := 0; ; count++ {
		where := fall(&sand, cave, height, true)

		if where == nil {
			return count
		} else if _, have := cave[*where]; have {
			return count
		}

		cave[*where] = 'o'
	}

	return 0
}

func sol2(cave map[Coordinate]int, height int) int {
	sand := Coordinate{500, 0}

	for count := 0; ; count++ {
		where := fall(&sand, cave, height+2, false)

		if where == nil {
			return count
		} else if _, have := cave[*where]; have {
			return count
		}

		cave[*where] = 'o'
	}

	return 0
}

func parse(lines []string) (map[Coordinate]int, int) {
	height := 0
	cave := map[Coordinate]int{}

	for _, line := range lines {
		rocks := strings.Split(line, "->")

		for i := 0; i < len(rocks)-1; i++ {
			head := atois(strings.Split(strings.TrimSpace(rocks[i]), ","))
			tail := atois(strings.Split(strings.TrimSpace(rocks[i+1]), ","))

			for x := min(head[0], tail[0]); x <= max(head[0], tail[0]); x++ {
				for y := min(head[1], tail[1]); y <= max(head[1], tail[1]); y++ {
					cave[Coordinate{x, y}] = '#'
				}
			}

			height = max(height, head[1])
			height = max(height, tail[1])
		}
	}

	return cave, height
}

func main() {
	lines := []string{}

	for input := range read(os.Stdin) {
		lines = append(lines, input)
	}

	fmt.Println("Ans1", sol1(parse(lines)))
	fmt.Println("Ans2", sol2(parse(lines)))
}
