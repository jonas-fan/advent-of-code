package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func read(reader io.Reader) chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		scanner := bufio.NewScanner(reader)

		for scanner.Scan() {
			out <- scanner.Text()
		}
	}()

	return out
}

func atoi(str string) int {
	out, _ := strconv.Atoi(str)

	return out
}

type Monkey struct {
	items     []int
	op        func(val int) int
	testVal   int
	testTrue  int
	testFalse int
}

func parseMonkey(lines []string) *Monkey {
	monkey := &Monkey{}
	tokens := strings.Split(lines[1], ":")

	for _, item := range strings.Split(tokens[1], ",") {
		monkey.items = append(monkey.items, atoi(strings.TrimSpace(item)))
	}

	tokens = strings.Split(lines[2], " ")
	ops := tokens[len(tokens)-2]
	val := atoi(tokens[len(tokens)-1])

	switch ops {
	case "+":
		monkey.op = func(rhs int) int {
			if val == 0 {
				return rhs << 1
			} else {
				return rhs + val
			}
		}
	case "*":
		monkey.op = func(rhs int) int {
			if val == 0 {
				return rhs * rhs
			} else {
				return rhs * val
			}
		}
	}

	tokens = strings.Split(lines[3], " ")
	monkey.testVal = atoi(tokens[len(tokens)-1])

	tokens = strings.Split(lines[4], " ")
	monkey.testTrue = atoi(tokens[len(tokens)-1])

	tokens = strings.Split(lines[5], " ")
	monkey.testFalse = atoi(tokens[len(tokens)-1])

	return monkey
}

func parseMonkeys(lines []string) []*Monkey {
	monkeys := []*Monkey{}
	buffers := []string{}

	for _, line := range lines {
		if line == "" {
			monkeys = append(monkeys, parseMonkey(buffers))
			buffers = buffers[:0]
		} else {
			buffers = append(buffers, line)
		}
	}

	return append(monkeys, parseMonkey(buffers))
}

func sol1(monkeys []*Monkey, rounds int) int {
	inspects := make([]int, len(monkeys))

	for ; rounds > 0; rounds-- {
		for i, monkey := range monkeys {
			inspects[i] += len(monkey.items)

			for _, item := range monkey.items {
				item = monkey.op(item) / 3

				if item%monkey.testVal == 0 {
					monkeys[monkey.testTrue].items = append(monkeys[monkey.testTrue].items, item)
				} else {
					monkeys[monkey.testFalse].items = append(monkeys[monkey.testFalse].items, item)
				}
			}

			monkey.items = monkey.items[:0]
		}
	}

	sort.Ints(inspects)

	return inspects[len(inspects)-1] * inspects[len(inspects)-2]
}

func sol2(monkeys []*Monkey, rounds int) int {
	inspects := make([]int, len(monkeys))
	boundary := 1

	for _, monkey := range monkeys {
		boundary *= monkey.testVal
	}

	for ; rounds > 0; rounds-- {
		for i, monkey := range monkeys {
			inspects[i] += len(monkey.items)

			for _, item := range monkey.items {
				item = monkey.op(item) % boundary

				if item%monkey.testVal == 0 {
					monkeys[monkey.testTrue].items = append(monkeys[monkey.testTrue].items, item)
				} else {
					monkeys[monkey.testFalse].items = append(monkeys[monkey.testFalse].items, item)
				}
			}
			monkey.items = monkey.items[:0]
		}
	}

	sort.Ints(inspects)

	return inspects[len(inspects)-1] * inspects[len(inspects)-2]
}

func main() {
	lines := []string{}

	for input := range read(os.Stdin) {
		lines = append(lines, input)
	}

	fmt.Println("Ans1", sol1(parseMonkeys(lines), 20))
	fmt.Println("Ans2", sol2(parseMonkeys(lines), 10000))
}
