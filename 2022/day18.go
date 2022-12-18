package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/jonas-fan/advent-of-code/funcs"
)

type Point struct {
	x int
	y int
	z int
}

func trapped(point Point, lava map[Point]bool, lower Point, upper Point) bool {
	var next Point

	for next = point; next.x >= lower.x; next.x-- {
		if _, have := lava[next]; have {
			break
		}
	}

	if next.x < lower.x {
		return false
	}

	for next = point; next.x <= upper.x; next.x++ {
		if _, have := lava[next]; have {
			break
		}
	}

	if next.x > upper.x {
		return false
	}

	for next = point; next.y >= lower.y; next.y-- {
		if _, have := lava[next]; have {
			break
		}
	}

	if next.y < lower.y {
		return false
	}

	for next = point; next.y <= upper.y; next.y++ {
		if _, have := lava[next]; have {
			break
		}
	}

	if next.y > upper.y {
		return false
	}

	for next = point; next.z >= lower.z; next.z-- {
		if _, have := lava[next]; have {
			break
		}
	}

	if next.z < lower.z {
		return false
	}

	for next = point; next.z <= upper.z; next.z++ {
		if _, have := lava[next]; have {
			break
		}
	}

	if next.z > upper.z {
		return false
	}

	return true
}

func sol(lava []Point, lower Point, upper Point, exterior bool) int {
	out := 0
	seen := map[Point]bool{}

	for _, point := range lava {
		seen[point] = true
	}

	deltas := []Point{
		Point{x: -1, y: 0, z: 0},
		Point{x: +1, y: 0, z: 0},
		Point{x: 0, y: -1, z: 0},
		Point{x: 0, y: +1, z: 0},
		Point{x: 0, y: 0, z: -1},
		Point{x: 0, y: 0, z: +1},
	}

	for _, point := range lava {
		out += 6

		for _, delta := range deltas {
			next := Point{
				x: point.x + delta.x,
				y: point.y + delta.y,
				z: point.z + delta.z,
			}

			if _, have := seen[next]; have {
				out--
			} else if exterior && trapped(next, seen, lower, upper) {
				out--
			}
		}
	}

	return out
}

func parse(lines []string) ([]Point, Point, Point) {
	lava := []Point{}

	lower := Point{
		x: math.MaxInt32,
		y: math.MaxInt32,
		z: math.MaxInt32,
	}

	upper := Point{
		x: 0,
		y: 0,
		z: 0,
	}

	for _, line := range lines {
		nums := funcs.Atois(strings.Split(line, ","))

		point := Point{
			x: nums[0],
			y: nums[1],
			z: nums[2],
		}

		lava = append(lava, point)

		lower.x = funcs.Min(lower.x, point.x)
		lower.y = funcs.Min(lower.y, point.y)
		lower.z = funcs.Min(lower.z, point.z)
		upper.x = funcs.Max(upper.x, point.x)
		upper.y = funcs.Max(upper.y, point.y)
		upper.z = funcs.Max(upper.z, point.z)
	}

	return lava, lower, upper
}

func main() {
	lines := []string{}

	for input := range funcs.ReadLines(os.Stdin) {
		lines = append(lines, input)
	}

	lava, lower, upper := parse(lines)

	fmt.Println("Ans1", sol(lava, lower, upper, false))
	fmt.Println("Ans2", sol(lava, lower, upper, true))
}
