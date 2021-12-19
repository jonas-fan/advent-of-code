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
	action := make(map[string]int)

	for _, each := range commands {
		tokens := strings.SplitN(each, " ", 2)

		val, _ := strconv.Atoi(tokens[1])

		action[tokens[0]] += val
	}

	return action["forward"] * (action["down"] - action["up"])
}

func main() {
	commands := make([]string, 0)

	for input := range read(os.Stdin) {
		commands = append(commands, input)
	}

	fmt.Println(solution(commands))
}
