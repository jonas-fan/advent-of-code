package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jonas-fan/advent-of-code/funcs"
)

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
			head := funcs.Atois(strings.Split(strings.TrimSpace(rocks[i]), ","))
			tail := funcs.Atois(strings.Split(strings.TrimSpace(rocks[i+1]), ","))

			for x := funcs.Min(head[0], tail[0]); x <= funcs.Max(head[0], tail[0]); x++ {
				for y := funcs.Min(head[1], tail[1]); y <= funcs.Max(head[1], tail[1]); y++ {
					cave[Coordinate{x, y}] = '#'
				}
			}

			height = funcs.Max(height, head[1], tail[1])
		}
	}

	return cave, height
}

func main() {
	lines := []string{}

	for input := range funcs.ReadLines(os.Stdin) {
		lines = append(lines, input)
	}

	fmt.Println("Ans1", sol1(parse(lines)))
	fmt.Println("Ans2", sol2(parse(lines)))
}
