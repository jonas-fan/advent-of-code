package main

import (
	"fmt"
	"os"

	"github.com/jonas-fan/advent-of-code/funcs"
)

type OreRobotSpec struct {
	ore int
}

type ClayRobotSpec struct {
	ore int
}

type ObsidianRobotSpec struct {
	ore  int
	clay int
}

type GeodeRobotSpec struct {
	ore      int
	obsidian int
}

var oreSpec OreRobotSpec
var claySpec ClayRobotSpec
var obsidianSpec ObsidianRobotSpec
var geodeSpec GeodeRobotSpec

type Args struct {
	times         int
	ore           int
	oreRobot      int
	clay          int
	clayRobot     int
	obsidian      int
	obsidianRobot int
	geode         int
	geodeRobot    int
}

func dfs(args Args, memo map[Args]int) int {
	key := args
	best := args.geode

	if args.times == 0 {
		return best
	} else if val, have := memo[key]; have {
		return val
	}

	args.times--

	if args.ore >= geodeSpec.ore && args.obsidian >= geodeSpec.obsidian {
		next := args

		next.ore = next.ore + next.oreRobot - geodeSpec.ore
		next.clay = next.clay + next.clayRobot
		next.obsidian = next.obsidian + next.obsidianRobot - geodeSpec.obsidian
		next.geode = next.geode + next.geodeRobot
		next.geodeRobot++

		best = funcs.Max(best, dfs(next, memo))
	} else if args.ore >= obsidianSpec.ore && args.clay >= obsidianSpec.clay {
		next := args

		next.ore = next.ore + next.oreRobot - obsidianSpec.ore
		next.clay = next.clay + next.clayRobot - obsidianSpec.clay
		next.obsidian = next.obsidian + next.obsidianRobot
		next.obsidianRobot++
		next.geode = next.geode + next.geodeRobot

		best = funcs.Max(best, dfs(next, memo))
	} else {
		if args.ore >= claySpec.ore {
			next := args

			next.ore = next.ore + next.oreRobot - claySpec.ore
			next.clay = next.clay + next.clayRobot
			next.clayRobot++
			next.obsidian = next.obsidian + next.obsidianRobot
			next.geode = next.geode + next.geodeRobot

			best = funcs.Max(best, dfs(next, memo))
		}

		if args.ore >= oreSpec.ore {
			next := args

			next.ore = next.ore + next.oreRobot - oreSpec.ore
			next.oreRobot++
			next.clay = next.clay + next.clayRobot
			next.obsidian = next.obsidian + next.obsidianRobot
			next.geode = next.geode + next.geodeRobot

			best = funcs.Max(best, dfs(next, memo))
		}

		next := args

		next.ore = next.ore + next.oreRobot
		next.clay = next.clay + next.clayRobot
		next.obsidian = next.obsidian + next.obsidianRobot
		next.geode = next.geode + next.geodeRobot

		best = funcs.Max(best, dfs(next, memo))
	}

	memo[key] = best

	return best
}

func sol1(lines []string, times int) int {
	out := 0
	round := 0
	format := "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian."

	for _, line := range lines {
		fmt.Sscanf(line, format,
			&round,
			&oreSpec.ore,
			&claySpec.ore,
			&obsidianSpec.ore, &obsidianSpec.clay,
			&geodeSpec.ore, &geodeSpec.obsidian)

		args := Args{
			times:    times,
			oreRobot: 1,
		}

		out += round * dfs(args, map[Args]int{})
	}

	return out
}

func sol2(lines []string, times int, n int) int {
	out := 1
	round := 0
	format := "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian."

	for i := 0; i < funcs.Min(len(lines), n); i++ {
		fmt.Sscanf(lines[i], format,
			&round,
			&oreSpec.ore,
			&claySpec.ore,
			&obsidianSpec.ore, &obsidianSpec.clay,
			&geodeSpec.ore, &geodeSpec.obsidian)

		args := Args{
			times:    times,
			oreRobot: 1,
		}

		out *= dfs(args, map[Args]int{})
	}

	return out
}

func main() {
	lines := []string{}

	for input := range funcs.ReadLines(os.Stdin) {
		lines = append(lines, input)
	}

	fmt.Println("Ans1", sol1(lines, 24))
	fmt.Println("Ans2", sol2(lines, 32, 3))
}
