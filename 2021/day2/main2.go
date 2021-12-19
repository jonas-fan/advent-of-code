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

func solution(commands []string) int {
	pos := 0
	aim := 0
	depth := 0

	for _, each := range commands {
		tokens := strings.SplitN(each, " ", 2)
		val, _ := strconv.Atoi(tokens[1])

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

	fmt.Println(solution(commands))
}
