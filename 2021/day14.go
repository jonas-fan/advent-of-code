package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func solution(template string, rules map[string]string, steps int) int {
	tokens := make(map[string]int)

	for i := 0; i < len(template)-1; i++ {
		tokens[template[i:i+2]]++
	}

	for ; steps > 0; steps-- {
		replaced := make(map[string]int)

		for token, count := range tokens {
			replacement := rules[token]

			replaced[string(token[0])+replacement] += count
			replaced[replacement+string(token[1])] += count
		}

		tokens = replaced
	}

	letter := map[byte]int{
		template[0]: 1,
	}

	for token, count := range tokens {
		letter[token[1]] += count
	}

	count := make([]int, 0, len(letter))

	for _, each := range letter {
		count = append(count, each)
	}

	sort.Slice(count, func(lhs int, rhs int) bool {
		return count[lhs] > count[rhs]
	})

	return count[0] - count[len(count)-1]
}

func main() {
	template := ""
	rules := make(map[string]string)
	reader := read(os.Stdin)

	for input := range reader {
		if input == "" {
			break
		}

		template = input
	}

	for input := range reader {
		tokens := strings.Split(input, " -> ")

		rules[tokens[0]] = tokens[1]
	}

	fmt.Println(solution(template, rules, 10))
	fmt.Println(solution(template, rules, 40))
}
