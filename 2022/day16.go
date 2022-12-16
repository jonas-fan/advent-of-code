package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/jonas-fan/advent-of-code/funcs"
)

type Room struct {
	flow    int
	tunnels []string
}

var initWhere string
var initTimes int

func dfs(rooms map[string]Room, rounds int, where string, times int,
	opened map[string]bool, memo map[string]int) int {
	if rounds <= 0 {
		return 0
	} else if times <= 0 {
		return dfs(rooms, rounds-1, initWhere, initTimes, opened, memo)
	}

	signature := fmt.Sprintf("%v%v%v%v", rounds, where, times, opened)

	if val, have := memo[signature]; have {
		return val
	}

	max := 0
	times--

	// try not to open
	for _, next := range rooms[where].tunnels {
		max = funcs.Max(max, dfs(rooms, rounds, next, times, opened, memo))
	}

	// try to open
	if _, have := opened[where]; !have {
		if rooms[where].flow > 0 {
			flow := rooms[where].flow * times

			opened[where] = true

			for _, next := range rooms[where].tunnels {
				max = funcs.Max(max, flow+dfs(rooms, rounds, next, times-1, opened, memo))
			}

			delete(opened, where)
		}
	}

	memo[signature] = max

	return max
}

func sol1(rooms map[string]Room) int {
	initWhere = "AA"
	initTimes = 30

	return dfs(rooms, 1, initWhere, initTimes, map[string]bool{}, map[string]int{})
}

func sol2(rooms map[string]Room) int {
	initWhere = "AA"
	initTimes = 26

	return dfs(rooms, 2, initWhere, initTimes, map[string]bool{}, map[string]int{})
}

func parse(lines []string) map[string]Room {
	rooms := map[string]Room{}
	re := regexp.MustCompile(`^Valve (\w+) has flow rate=(\d+); tunnel[s]* lead[s]* to valve[s]* (.*)$`)

	for _, line := range lines {
		matches := re.FindStringSubmatch(line)

		rooms[matches[1]] = Room{
			flow:    funcs.Atoi(matches[2]),
			tunnels: strings.Split(matches[3], ", "),
		}
	}

	return rooms
}

func main() {
	lines := []string{}

	for input := range funcs.ReadLines(os.Stdin) {
		lines = append(lines, input)
	}

	fmt.Println("Ans1", sol1(parse(lines)))
	fmt.Println("Ans2", sol2(parse(lines)))
}
