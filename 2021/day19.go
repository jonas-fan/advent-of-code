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

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
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

type Coordinate struct {
	x int
	y int
	z int
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("(%d, %d, %d)", c.x, c.y, c.z)
}

func (c *Coordinate) Equal(other *Coordinate) bool {
	if c.x != other.x {
		return false
	} else if c.y != other.y {
		return false
	} else if c.z != other.z {
		return false
	}

	return true
}

func (c *Coordinate) Rotate(f int, r int) *Coordinate {
	out := &Coordinate{
		x: c.x,
		y: c.y,
		z: c.z,
	}

	switch f {
	case 0:
		out.x, out.y, out.z = out.x, out.y, out.z
	case 1:
		out.x, out.y, out.z = -out.z, out.y, out.x
	case 2:
		out.x, out.y, out.z = -out.x, out.y, -out.z
	case 3:
		out.x, out.y, out.z = out.z, out.y, -out.x
	case 4:
		out.x, out.y, out.z = out.y, -out.x, out.z
	case 5:
		out.x, out.y, out.z = -out.y, out.x, out.z
	}

	switch r {
	case 0:
		out.x, out.y, out.z = out.x, out.y, out.z
	case 1:
		out.x, out.y, out.z = out.x, out.z, -out.y
	case 2:
		out.x, out.y, out.z = out.x, -out.y, -out.z
	case 3:
		out.x, out.y, out.z = out.x, -out.z, out.y
	}

	return out
}

func rotate(beacons []*Coordinate, f int, r int) []*Coordinate {
	out := make([]*Coordinate, 0, len(beacons))

	for _, each := range beacons {
		out = append(out, each.Rotate(f, r))
	}

	return out
}

func relative(positions []*Coordinate) [][]*Coordinate {
	out := make([][]*Coordinate, 0, len(positions))

	for _, lhs := range positions {
		diff := make([]*Coordinate, 0, len(positions))

		for _, rhs := range positions {
			diff = append(diff, &Coordinate{
				x: rhs.x - lhs.x,
				y: rhs.y - lhs.y,
				z: rhs.z - lhs.z,
			})
		}

		out = append(out, diff)
	}

	return out
}

func overlapped(lhs []*Coordinate, rhs []*Coordinate) map[int]int {
	laps := make(map[int]int)
	ldiff := relative(lhs)
	rdiff := relative(rhs)

	for lfrom := range ldiff {
		for lto := range ldiff[lfrom] {
			if lfrom == lto {
				continue
			}

			for rfrom := range rdiff {
				for rto := range rdiff[rfrom] {
					if rfrom == rto {
						continue
					}

					if ldiff[lfrom][lto].Equal(rdiff[rfrom][rto]) {
						laps[lfrom] = rfrom
						laps[lto] = rto
					}
				}
			}
		}
	}

	return laps
}

func findOverlapped(lhs []*Coordinate, rhs []*Coordinate) (map[int]int, []*Coordinate) {
	for f := 0; f < 6; f++ {
		for r := 0; r < 4; r++ {
			rotated := rotate(rhs, f, r)
			laps := overlapped(lhs, rotated)

			if len(laps) == 12 {
				return laps, rotated
			}
		}
	}

	return nil, nil
}

func solution(scanners [][]*Coordinate) (int, int) {
	positions := make([]*Coordinate, len(scanners))

	positions[0] = &Coordinate{}

	for n := 0; n < len(scanners); n++ {
		for i := 0; i < len(scanners); i++ {
			for j := 0; j < len(scanners); j++ {
				if positions[i] == nil || positions[j] != nil {
					continue
				}

				laps, rotated := findOverlapped(scanners[i], scanners[j])

				if laps != nil {
					scanners[j] = rotated

					for l, r := range laps {
						positions[j] = &Coordinate{
							x: scanners[i][l].x - scanners[j][r].x + positions[i].x,
							y: scanners[i][l].y - scanners[j][r].y + positions[i].y,
							z: scanners[i][l].z - scanners[j][r].z + positions[i].z,
						}

						break
					}
				}
			}
		}
	}

	coordinateMap := make(map[Coordinate]bool)

	for i, position := range positions {
		for _, each := range scanners[i] {
			coordinate := Coordinate{
				x: position.x + each.x,
				y: position.y + each.y,
				z: position.z + each.z,
			}

			coordinateMap[coordinate] = true
		}
	}

	maxDistance := 0

	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			from := positions[i]
			to := positions[j]

			distance := abs(from.x-to.x) + abs(from.y-to.y) + abs(from.z-to.z)

			maxDistance = max(maxDistance, distance)
		}
	}

	return len(coordinateMap), maxDistance
}

func main() {
	scanners := [][]*Coordinate{}
	reader := read(os.Stdin)
	re := regexp.MustCompile(`--- scanner ([0-9]+) ---`)

	for input := range reader {
		matches := re.FindStringSubmatch(input)

		if len(matches) < 2 {
			continue
		}

		beacons := []*Coordinate{}

		for input := range reader {
			tokens := str2intSlice(strings.Split(input, ","))

			if len(tokens) != 3 {
				break
			}

			beacons = append(beacons, &Coordinate{
				x: tokens[0],
				y: tokens[1],
				z: tokens[2],
			})
		}

		scanners = append(scanners, beacons)
	}

	fmt.Println(solution(scanners))
}
