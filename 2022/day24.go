package main

import (
	"fmt"
	"os"

	"github.com/jonas-fan/advent-of-code/funcs"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type Point struct {
	x int
	y int
}

type Blizzard struct {
	Point
	direct int
}

func move(blizzard []*Blizzard, width int, height int) []*Blizzard {
	for _, b := range blizzard {
		switch b.direct {
		case UP:
			b.y--
		case DOWN:
			b.y++
		case LEFT:
			b.x--
		case RIGHT:
			b.x++
		}

		b.x = (b.x + width) % width
		b.y = (b.y + height) % height
	}

	return blizzard
}

func sol1(blizzard []*Blizzard, width int, height int, begin Point, end Point) int {
	m := map[Point][]*Blizzard{}

	for _, b := range blizzard {
		m[b.Point] = append(m[b.Point], b)
	}

	queue := []Point{begin}

	for round := 0; ; round++ {
		draft := map[Point][]*Blizzard{}

		for _, b := range move(blizzard, width, height) {
			draft[b.Point] = append(draft[b.Point], b)
		}

		avail := map[Point]bool{}

		for _, point := range queue {
			next := []Point{
				Point{point.x - 1, point.y},
				Point{point.x + 1, point.y},
				Point{point.x, point.y - 1},
				Point{point.x, point.y + 1},
			}

			for _, each := range next {
				if each.x < 0 || each.x >= width {
					continue
				} else if each.y < 0 || each.y >= height {
					continue
				} else if _, have := draft[each]; have {
					continue
				}

				avail[each] = true
			}

			if _, have := draft[point]; !have {
				avail[point] = true
			}
		}

		queue = queue[:0]

		for point := range avail {
			if funcs.Abs(point.x-end.x)+funcs.Abs(point.y-end.y) < 2 {
				move(blizzard, width, height)
				return round + 2
			}

			queue = append(queue, point)
		}

		m = draft
	}

	return 0
}

func makeBlizzard(lines []string) []*Blizzard {
	out := []*Blizzard{}

	for y, line := range lines {
		for x, letter := range line {
			b := &Blizzard{
				Point: Point{x: x - 1, y: y - 1},
			}

			if letter == '^' {
				b.direct = UP
			} else if letter == 'v' {
				b.direct = DOWN
			} else if letter == '<' {
				b.direct = LEFT
			} else if letter == '>' {
				b.direct = RIGHT
			} else {
				continue
			}

			out = append(out, b)
		}
	}

	return out
}

func sol2(blizzard []*Blizzard, width int, height int, begin Point, end Point) int {
	out := 0

	out += sol1(blizzard, width, height, begin, end)
	out += sol1(blizzard, width, height, end, begin)
	out += sol1(blizzard, width, height, begin, end)

	return out
}

func main() {
	lines := []string{}

	for input := range funcs.ReadLines(os.Stdin) {
		lines = append(lines, input)
	}

	width := len(lines[0]) - 2
	height := len(lines) - 2

	begin := Point{x: 0, y: -1}
	end := Point{x: width - 1, y: height}

	fmt.Println("Ans1", sol1(makeBlizzard(lines), width, height, begin, end))
	fmt.Println("Ans2", sol2(makeBlizzard(lines), width, height, begin, end))
}
