package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
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

func parseStacks(lines []string) [][]rune {
	stacks := [][]rune{}

	for col := 1; col < len(lines[0]); col += 4 {
		stack := []rune{}

		for row := len(lines) - 2; row >= 0; row-- {
			letter := rune(lines[row][col])

			if unicode.IsLetter(letter) {
				stack = append(stack, letter)
			}
		}

		stacks = append(stacks, stack)
	}

	return stacks
}

func parseStep(line string) (loop int, from int, to int) {
	fmt.Sscanf(line, "move %d from %d to %d", &loop, &from, &to)

	from--
	to--

	return
}

func sol1(stacks [][]rune, lines []string) string {
	for _, line := range lines {
		loop, from, to := parseStep(line)

		for i := 0; i < loop; i++ {
			fromTop := len(stacks[from]) - 1
			what := stacks[from][fromTop]

			stacks[from] = stacks[from][:fromTop]
			stacks[to] = append(stacks[to], what)
		}
	}

	out := []rune{}

	for _, stack := range stacks {
		out = append(out, stack[len(stack)-1])
	}

	return string(out)
}

func sol2(stacks [][]rune, lines []string) string {
	for _, line := range lines {
		loop, from, to := parseStep(line)
		fromTop := len(stacks[from]) - loop
		whats := stacks[from][fromTop:]

		stacks[from] = stacks[from][:fromTop]
		stacks[to] = append(stacks[to], whats...)
	}

	out := []rune{}

	for _, stack := range stacks {
		out = append(out, stack[len(stack)-1])
	}

	return string(out)
}

func main() {
	stacks := []string{}
	lines := []string{}

	for input := range read(os.Stdin) {
		lines = append(lines, input)

		if input == "" {
			stacks = append(stacks, lines...)
			lines = lines[:0]
		}
	}

	fmt.Println("Ans1", sol1(parseStacks(stacks), lines))
	fmt.Println("Ans2", sol2(parseStacks(stacks), lines))
}
