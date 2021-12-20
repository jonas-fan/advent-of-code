package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
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

func isSmallCave(pos string) bool {
	for _, each := range pos {
		return unicode.IsLower(each)
	}

	return false
}

func explore1(pathmap map[string][]string, pos string, visited map[string]int) int {
	if pos == "end" {
		return 1
	}

	if isSmallCave(pos) {
		visited[pos]++
	}

	places := []string{}

	for _, next := range pathmap[pos] {
		switch {
		case pos == "start":
			places = append(places, next)
		case next == "start":
			// ignore
		case next == "end":
			places = append(places, next)
		case isSmallCave(next) && visited[next] > 0:
			// ignore
		default:
			places = append(places, next)
		}
	}

	sum := 0

	for _, next := range places {
		sum += explore1(pathmap, next, visited)
	}

	if isSmallCave(pos) {
		visited[pos]--

	}

	return sum
}

func solution1(pathmap map[string][]string) int {
	return explore1(pathmap, "start", map[string]int{})
}

func explore2(pathmap map[string][]string, pos string, visited map[string]int) int {
	if pos == "end" {
		return 1
	}

	if isSmallCave(pos) {
		visited[pos]++
	}

	maxVisits := 0

	for _, each := range visited {
		if each > maxVisits {
			maxVisits = each
		}
	}

	places := []string{}

	for _, next := range pathmap[pos] {
		switch {
		case pos == "start":
			places = append(places, next)
		case next == "start":
			// ignore
		case next == "end":
			places = append(places, next)
		case isSmallCave(next) && (maxVisits > 1 && visited[next] > 0):
			// ignore
		default:
			places = append(places, next)
		}
	}

	sum := 0

	for _, next := range places {
		sum += explore2(pathmap, next, visited)
	}

	if isSmallCave(pos) {
		visited[pos]--
	}

	return sum
}

func solution2(pathmap map[string][]string) int {
	return explore2(pathmap, "start", map[string]int{})
}

func main() {
	pathmap := make(map[string][]string)

	for input := range read(os.Stdin) {
		path := strings.Split(input, "-")

		pathmap[path[0]] = append(pathmap[path[0]], path[1])
		pathmap[path[1]] = append(pathmap[path[1]], path[0])
	}

	fmt.Println(solution1(pathmap))
	fmt.Println(solution2(pathmap))
}
