package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/jonas-fan/advent-of-code/funcs"
)

type Point struct {
	x int
	y int
}

func dist(lhs *Point, rhs *Point) int {
	return funcs.Abs(lhs.x-rhs.x) + funcs.Abs(lhs.y-rhs.y)
}

type Bound struct {
	left  int
	right int
}

func consecutive(bounds []*Bound) (bool, int) {
	upper := 0

	for i := 0; i < len(bounds)-1; i++ {
		upper = funcs.Max(upper, bounds[i].right)

		if upper+1 < bounds[i+1].left {
			return false, upper + 1
		}

		upper = funcs.Max(upper, bounds[i+1].right)
	}

	return true, upper + 1
}

func sol1(sensors []*Point, beacons []*Point, targetY int) int {
	// figure out the coverage of each row
	covered := map[Point]int{}

	for i := range sensors {
		sensor := sensors[i]
		beacon := beacons[i]
		distance := dist(sensor, beacon)

		for y := sensor.y - distance; y <= sensor.y+distance; y++ {
			if y != targetY {
				continue
			}

			rest := distance - funcs.Abs(y-sensor.y)

			for x := sensor.x - rest; x <= sensor.x+rest; x++ {
				covered[Point{x, y}] = '#'
			}
		}
	}

	// put the sensors and beacons
	for i := range sensors {
		covered[*sensors[i]] = 'S'
		covered[*beacons[i]] = 'B'
	}

	// find out the covered positions in the area
	out := 0

	for point, val := range covered {
		if point.y == targetY {
			if val == '#' {
				out++
			}
		}
	}

	return out
}

func sol2(sensors []*Point, beacons []*Point) int {
	// figure out the coverage of each row
	covered := map[int][]*Bound{}

	for i := range sensors {
		sensor := sensors[i]
		beacon := beacons[i]
		distance := dist(sensor, beacon)

		for y := sensor.y - distance; y <= sensor.y+distance; y++ {
			rest := distance - funcs.Abs(y-sensor.y)
			bound := &Bound{
				left:  sensor.x - rest,
				right: sensor.x + rest,
			}

			covered[y] = append(covered[y], bound)
		}
	}

	// find out the missing position in the area
	inside := false

	for y := 0; y < 4000000; y++ {
		bounds, have := covered[y]

		if !have {
			continue
		}

		sort.Slice(bounds, func(i int, j int) bool {
			return bounds[i].left < bounds[j].left
		})

		if inside {
			if ok, missingX := consecutive(bounds); !ok {
				return missingX*4000000 + y
			}
		} else {
			inside, _ = consecutive(bounds)
		}
	}

	return 0
}

func parse(lines []string) (sensors []*Point, beacons []*Point) {
	sensors = make([]*Point, 0, len(lines))
	beacons = make([]*Point, 0, len(lines))

	for _, line := range lines {
		sensor := &Point{}
		beacon := &Point{}

		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.x, &sensor.y, &beacon.x, &beacon.y)

		sensors = append(sensors, sensor)
		beacons = append(beacons, beacon)
	}

	return
}

func main() {
	lines := []string{}

	for input := range funcs.ReadLines(os.Stdin) {
		lines = append(lines, input)
	}

	sensors, beacons := parse(lines)

	fmt.Println("Ans1", sol1(sensors, beacons, 2000000))
	fmt.Println("Ans2", sol2(sensors, beacons))
}
