package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}

	return lhs
}

func sol1(line string) int {
	begin := -1
	seen := map[byte]int{}

	for i := 0; i < len(line); i++ {
		if where, have := seen[line[i]]; have {
			begin = max(begin, where)
		}

		seen[line[i]] = i

		if i-begin > 3 {
			return i + 1
		}
	}

	return 0
}

func sol2(line string) int {
	begin := -1
	seen := map[byte]int{}

	for i := 0; i < len(line); i++ {
		if where, have := seen[line[i]]; have {
			begin = max(begin, where)
		}

		seen[line[i]] = i

		if i-begin > 13 {
			return i + 1
		}
	}

	return 0
}

func main() {
	lines := []string{}

	for input := range read(os.Stdin) {
		lines = append(lines, input)
	}

	fmt.Println("Ans1", sol1(lines[0]))
	fmt.Println("Ans2", sol2(lines[0]))
}
