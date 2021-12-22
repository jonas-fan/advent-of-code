package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func solution1(commands []string) int {
	action := make(map[string]int)

	for _, each := range commands {
		tokens := strings.SplitN(each, " ", 2)

		action[tokens[0]] += str2int(tokens[1])
	}

	return action["forward"] * (action["down"] - action["up"])
}

func solution2(commands []string) int {
	pos := 0
	aim := 0
	depth := 0

	for _, each := range commands {
		tokens := strings.SplitN(each, " ", 2)
		val := str2int(tokens[1])

		switch tokens[0] {
		case "down":
			aim += val
		case "up":
			aim -= val
		case "forward":
			pos += val
			depth += aim * val
		}
	}

	return pos * depth
}

func main() {
	commands := make([]string, 0)

	for input := range read(os.Stdin) {
		commands = append(commands, input)
	}

	fmt.Println(solution1(commands))
	fmt.Println(solution2(commands))
}
