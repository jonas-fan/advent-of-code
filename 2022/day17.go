package main

import (
	"fmt"
	"os"

	"github.com/jonas-fan/advent-of-code/funcs"
)

type Point struct {
	x int
	y int
}

func (p *Point) Move(direct byte) *Point {
	switch direct {
	case '>':
		return &Point{p.x + 1, p.y}
	case '<':
		return &Point{p.x - 1, p.y}
	case 'v':
		return &Point{p.x, p.y - 1}
	}

	panic("Unknown direction")
}

type Rock struct {
	points []Point
}

func (r *Rock) Copy() *Rock {
	out := &Rock{}

	for _, point := range r.points {
		out.points = append(out.points, point)
	}

	return out
}

func (r *Rock) Move(direct byte, cave map[Point]bool, dryRun bool) bool {
	destinations := []*Point{}

	for _, point := range r.points {
		next := point.Move(direct)

		if next.x < 0 || next.x > 6 {
			return false
		} else if next.y < 0 {
			return false
		} else if _, have := cave[*next]; have {
			return false
		}

		destinations = append(destinations, next)
	}

	if !dryRun {
		for i := range destinations {
			r.points[i] = *destinations[i]
		}
	}

	return true
}

func topPart(cave map[Point]bool, top int, n int) map[Point]bool {
	out := map[Point]bool{}

	for point := range cave {
		point.y = top - point.y

		if point.y >= 0 && point.y < n {
			out[point] = true
		}
	}

	return out
}

type Pair struct {
	arg1 int
	arg2 int
}

func sol(rocks []Rock, directs string, target int) int {
	var cave = map[Point]bool{}
	var caveTop = -1
	var rockCount int
	var directIndex int
	var memo = map[string]Pair{}

	for rockCount < target {
		rock := rocks[rockCount%len(rocks)].Copy()

		// update offset
		for i := range rock.points {
			rock.points[i].x += 2
			rock.points[i].y += caveTop + 4
		}

		// fall until stopped
		for {
			direct := directs[directIndex%len(directs)]

			directIndex++

			// left or right
			if rock.Move(direct, cave, true) {
				rock.Move(direct, cave, false)
			}

			// down
			if rock.Move('v', cave, true) {
				rock.Move('v', cave, false)
			} else {
				for _, point := range rock.points {
					cave[point] = true
					caveTop = funcs.Max(caveTop, point.y)
				}

				parts := topPart(cave, caveTop, 100)
				state := fmt.Sprint(rockCount%len(rocks), directIndex%len(directs), parts)

				if pair, have := memo[state]; have {
					rockIncreased := rockCount - pair.arg1
					caveTopIncreased := caveTop - pair.arg2
					times := (target - rockCount) / rockIncreased
					newCave := map[Point]bool{}

					for point := range cave {
						point.y += times * caveTopIncreased
						newCave[point] = true
					}

					cave = newCave
					rockCount += times * rockIncreased
					caveTop += times * caveTopIncreased
				}

				memo[state] = Pair{
					rockCount,
					caveTop,
				}

				break
			}
		}

		rockCount++
	}

	return caveTop + 1
}

func main() {
	lines := []string{}

	for input := range funcs.ReadLines(os.Stdin) {
		lines = append(lines, input)
	}

	rocks := []Rock{
		Rock{
			points: []Point{
				Point{0, 0},
				Point{1, 0},
				Point{2, 0},
				Point{3, 0},
			},
		},
		Rock{
			points: []Point{
				Point{1, 0},
				Point{0, 1},
				Point{1, 1},
				Point{2, 1},
				Point{1, 2},
			},
		},
		Rock{
			points: []Point{
				Point{0, 0},
				Point{1, 0},
				Point{2, 0},
				Point{2, 1},
				Point{2, 2},
			},
		},
		Rock{
			points: []Point{
				Point{0, 0},
				Point{0, 1},
				Point{0, 2},
				Point{0, 3},
			},
		},
		Rock{
			points: []Point{
				Point{0, 0},
				Point{0, 1},
				Point{1, 0},
				Point{1, 1},
			},
		},
	}

	fmt.Println("Ans1", sol(rocks, lines[0], 2022))
	fmt.Println("Ans2", sol(rocks, lines[0], 1000000000000))
}
