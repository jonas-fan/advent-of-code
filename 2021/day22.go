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

func str2int(str string) int {
	val, _ := strconv.Atoi(str)

	return val
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

type Range struct {
	From int
	To   int
}

type Cube struct {
	On bool
	X  Range
	Y  Range
	Z  Range
}

func overlapped(lhs *Cube, rhs *Cube) (bool, *Cube) {
	out := &Cube{
		On: !lhs.On,
		X:  Range{From: max(lhs.X.From, rhs.X.From), To: min(lhs.X.To, rhs.X.To)},
		Y:  Range{From: max(lhs.Y.From, rhs.Y.From), To: min(lhs.Y.To, rhs.Y.To)},
		Z:  Range{From: max(lhs.Z.From, rhs.Z.From), To: min(lhs.Z.To, rhs.Z.To)},
	}

	switch {
	case out.X.From > out.X.To:
		return false, nil
	case out.Y.From > out.Y.To:
		return false, nil
	case out.Z.From > out.Z.To:
		return false, nil
	}

	return true, out
}

func solution(cubes []Cube, size int) int {
	refactored := []Cube{}

	for _, cube := range cubes {
		if size > 0 {
			if cube.X.From > size || cube.X.To < -size {
				continue
			} else if cube.Y.From > size || cube.Y.To < -size {
				continue
			} else if cube.Z.From > size || cube.Z.To < -size {
				continue
			}

			cube.X.From = max(cube.X.From, -size)
			cube.X.To = min(cube.X.To, size)
			cube.Y.From = max(cube.Y.From, -size)
			cube.Y.To = min(cube.Y.To, size)
			cube.Z.From = max(cube.Z.From, -size)
			cube.Z.To = min(cube.Z.To, size)
		}

		frames := []Cube{}

		for _, c := range refactored {
			if ok, overlap := overlapped(&c, &cube); ok {
				frames = append(frames, *overlap)
			}
		}

		if cube.On {
			frames = append(frames, cube)
		}

		refactored = append(refactored, frames...)
	}

	lights := 0

	for _, cube := range refactored {
		count := (cube.X.To - cube.X.From + 1) *
			(cube.Y.To - cube.Y.From + 1) *
			(cube.Z.To - cube.Z.From + 1)

		if cube.On {
			lights += count
		} else {
			lights -= count
		}
	}

	return lights
}

func main() {
	pattern := regexp.MustCompile(`([^ ]+) x=([^.]+)..([^.]+),y=([^.]+)..([^.]+),z=([^.]+)..([^.]+)`)
	cubes := []Cube{}

	for input := range read(os.Stdin) {
		matches := pattern.FindStringSubmatch(input)

		if len(matches) < 8 {
			continue
		}

		cube := Cube{
			On: matches[1] == "on",
			X: Range{
				From: min(str2int(matches[2]), str2int(matches[3])),
				To:   max(str2int(matches[2]), str2int(matches[3])),
			},
			Y: Range{
				From: min(str2int(matches[4]), str2int(matches[5])),
				To:   max(str2int(matches[4]), str2int(matches[5])),
			},
			Z: Range{
				From: min(str2int(matches[6]), str2int(matches[7])),
				To:   max(str2int(matches[6]), str2int(matches[7])),
			},
		}

		cubes = append(cubes, cube)
	}

	fmt.Println(solution(cubes, 50))
	fmt.Println(solution(cubes, 0))
}
