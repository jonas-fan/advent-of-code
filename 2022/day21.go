package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jonas-fan/advent-of-code/funcs"
)

func dfs1(m map[string]string, name string) int {
	val := m[name]

	if num, err := strconv.Atoi(val); err == nil {
		return num
	}

	tokens := strings.Split(val, " ")
	lhs := tokens[0]
	ops := tokens[1]
	rhs := tokens[2]

	switch ops {
	case "+":
		return dfs1(m, lhs) + dfs1(m, rhs)
	case "-":
		return dfs1(m, lhs) - dfs1(m, rhs)
	case "*":
		return dfs1(m, lhs) * dfs1(m, rhs)
	case "/":
		return dfs1(m, lhs) / dfs1(m, rhs)
	}

	panic("unknown ops")

	return 0
}

func sol1(m map[string]string) int {
	return dfs1(m, "root")
}

func canResolve(m map[string]string, name string) bool {
	if name == "humn" {
		return false
	}

	val := m[name]

	if _, err := strconv.Atoi(val); err == nil {
		return true
	}

	tokens := strings.Split(val, " ")
	lhs := tokens[0]
	rhs := tokens[2]

	return canResolve(m, lhs) && canResolve(m, rhs)
}

func dfs2(m map[string]string, name string, target int) int {
	val := m[name]

	if _, err := strconv.Atoi(val); err == nil {
		return target
	}

	tokens := strings.Split(val, " ")
	lhs := tokens[0]
	ops := tokens[1]
	rhs := tokens[2]

	if name == "root" {
		ops = "="
	}

	if canResolve(m, lhs) {
		lvalue := dfs1(m, lhs)

		switch ops {
		case "=":
			return dfs2(m, rhs, lvalue)
		case "+":
			return dfs2(m, rhs, target-lvalue)
		case "-":
			return dfs2(m, rhs, lvalue-target)
		case "*":
			return dfs2(m, rhs, target/lvalue)
		case "/":
			return dfs2(m, rhs, lvalue/target)
		}
	} else if canResolve(m, rhs) {
		rvalue := dfs1(m, rhs)

		switch ops {
		case "=":
			return dfs2(m, lhs, rvalue)
		case "+":
			return dfs2(m, lhs, target-rvalue)
		case "-":
			return dfs2(m, lhs, target+rvalue)
		case "*":
			return dfs2(m, lhs, target/rvalue)
		case "/":
			return dfs2(m, lhs, target*rvalue)
		}
	}

	panic("cannot resolve")

	return 0
}

func sol2(m map[string]string) int {
	return dfs2(m, "root", 0)
}

func main() {
	m := map[string]string{}

	for input := range funcs.ReadLines(os.Stdin) {
		tokens := strings.Split(input, ": ")

		m[tokens[0]] = tokens[1]
	}

	fmt.Println("Ans1", sol1(m))
	fmt.Println("Ans2", sol2(m))
}
