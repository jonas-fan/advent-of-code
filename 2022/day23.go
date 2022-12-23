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

type Elf struct {
	Point
	direct int
}

func (e *Elf) Alone(m map[Point]bool) bool {
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if x == 0 && y == 0 {
				continue
			}

			if _, have := m[Point{e.x + x, e.y + y}]; have {
				return false
			}
		}
	}

	return true
}

func (e *Elf) CanMoveUp(m map[Point]bool) bool {
	for offset := -1; offset < 2; offset++ {
		if _, have := m[Point{e.x + offset, e.y - 1}]; have {
			return false
		}
	}

	return true
}

func (e *Elf) CanMoveDown(m map[Point]bool) bool {
	for offset := -1; offset < 2; offset++ {
		if _, have := m[Point{e.x + offset, e.y + 1}]; have {
			return false
		}
	}

	return true
}

func (e *Elf) CanMoveLeft(m map[Point]bool) bool {
	for offset := -1; offset < 2; offset++ {
		if _, have := m[Point{e.x - 1, e.y + offset}]; have {
			return false
		}
	}

	return true
}

func (e *Elf) CanMoveRight(m map[Point]bool) bool {
	for offset := -1; offset < 2; offset++ {
		if _, have := m[Point{e.x + 1, e.y + offset}]; have {
			return false
		}
	}

	return true
}

func same(lhs map[Point]bool, rhs map[Point]bool) bool {
	if len(lhs) != len(rhs) {
		return false
	}

	for lkey, lvalue := range lhs {
		if rvalue, have := rhs[lkey]; !have {
			return false
		} else if rvalue != lvalue {
			return false
		}
	}

	return true
}

func sol(elves []*Elf, rounds int) int {
	m := map[Point]bool{}

	for _, e := range elves {
		m[e.Point] = true
	}

	round := 0

	for ; round < rounds; round++ {
		draft := map[Point][]*Elf{}

		for _, elf := range elves {
			if !elf.Alone(m) {
				for r := round; r < round+4; r++ {
					elf.direct = r % 4

					if elf.direct == UP && elf.CanMoveUp(m) {
						elf.y--
						break
					} else if elf.direct == DOWN && elf.CanMoveDown(m) {
						elf.y++
						break
					} else if elf.direct == LEFT && elf.CanMoveLeft(m) {
						elf.x--
						break
					} else if elf.direct == RIGHT && elf.CanMoveRight(m) {
						elf.x++
						break
					}
				}
			}

			draft[elf.Point] = append(draft[elf.Point], elf)
		}

		next := map[Point]bool{}

		for _, elves := range draft {
			if len(elves) == 1 {
				// move
				for _, elf := range elves {
					next[elf.Point] = true
				}
			} else {
				// conflicts, go back :(
				for _, elf := range elves {
					switch elf.direct {
					case UP:
						elf.y++
					case DOWN:
						elf.y--
					case LEFT:
						elf.x++
					case RIGHT:
						elf.x--
					}

					next[elf.Point] = true
				}
			}
		}

		if same(next, m) {
			round++
			break
		}

		m = next
	}

	fmt.Println("End Round:", round)

	minx, maxx := elves[0].x, elves[0].x
	miny, maxy := elves[0].y, elves[0].y

	for p := range m {
		minx = funcs.Min(minx, p.x)
		maxx = funcs.Max(maxx, p.x)
		miny = funcs.Min(miny, p.y)
		maxy = funcs.Max(maxy, p.y)
	}

	return (maxx-minx+1)*(maxy-miny+1) - len(m)
}

func makeElves(points []Point) []*Elf {
	out := []*Elf{}

	for _, point := range points {
		out = append(out, &Elf{Point: point})
	}

	return out
}

func main() {
	points := []Point{}
	y := 0

	for input := range funcs.ReadLines(os.Stdin) {
		for x, letter := range input {
			if letter == '#' {
				points = append(points, Point{x: x, y: y})
			}
		}

		y++
	}

	fmt.Println("Ans1", sol(makeElves(points), 10))
	fmt.Println("Ans2", sol(makeElves(points), 1000000000))
}
