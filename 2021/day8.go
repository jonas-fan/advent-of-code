package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

type RuneSlice []rune

func (r RuneSlice) Len() int {
	return len(r)
}

func (r RuneSlice) Less(lhs int, rhs int) bool {
	return r[lhs] < r[rhs]
}

func (r RuneSlice) Swap(lhs int, rhs int) {
	r[lhs], r[rhs] = r[rhs], r[lhs]
}

func filter(lhs string, rhs string) string {
	out := []rune{}

	for _, lvalue := range lhs {
		for _, rvalue := range rhs {
			if lvalue == rvalue {
				out = append(out, lvalue)
			}
		}
	}

	return string(out)
}

func not(bits string) string {
	out := []rune("abcdefg")

	for _, each := range bits {
		for i := range out {
			if out[i] == each {
				out = append(out[:i], out[i+1:]...)
				break
			}
		}
	}

	return string(out)
}

func and(lhs string, rhs string) bool {
	for _, lvalue := range lhs {
		for _, rvalue := range rhs {
			if lvalue == rvalue {
				return true
			}
		}
	}

	return false
}

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

func solution1(entries [][]string) int {
	tokens := []RuneSlice{}

	for _, entry := range entries {
		for _, pattern := range strings.Split(entry[1], " ") {
			token := RuneSlice(pattern)

			sort.Sort(token)

			tokens = append(tokens, token)
		}
	}

	count := 0

	for _, token := range tokens {
		switch len(token) {
		case 2, 3, 4, 7:
			count++
		}
	}

	return count
}

func solution2(entries [][]string) int {
	sum := 0

	for _, entry := range entries {
		tokens := []string{}

		for _, pattern := range strings.Split(entry[0], " ") {
			token := RuneSlice(pattern)

			sort.Sort(token)

			tokens = append(tokens, string(token))
		}

		digit := make(map[string]int)
		wire := make([]string, 10)

		for _, token := range tokens {
			switch len(token) {
			case 2:
				digit[token], wire[1] = 1, token
			case 3:
				digit[token], wire[7] = 7, token
			case 4:
				digit[token], wire[4] = 4, token
			case 7:
				digit[token], wire[8] = 8, token
			}
		}

		for _, token := range tokens {
			switch len(token) {
			case 5:
				switch {
				case !and(not(token), wire[1]):
					digit[token], wire[3] = 3, token
				case len(filter(token, wire[4])) == 2:
					digit[token], wire[2] = 2, token
				default:
					digit[token], wire[5] = 5, token
				}
			case 6:
				switch {
				case and(not(token), wire[1]) && and(not(token), wire[4]):
					digit[token], wire[6] = 6, token
				case !and(not(token), wire[1]) && !and(not(token), wire[4]):
					digit[token], wire[9] = 9, token
				default:
					digit[token], wire[0] = 0, token
				}
			}
		}

		num := 0

		for _, pattern := range strings.Split(entry[1], " ") {
			token := RuneSlice(pattern)

			sort.Sort(token)

			num = num*10 + digit[string(token)]
		}

		sum += num
	}

	return sum
}

func main() {
	entries := [][]string{}

	for input := range read(os.Stdin) {
		entry := strings.Split(input, " | ")

		entries = append(entries, entry)
	}

	fmt.Println(solution1(entries))
	fmt.Println(solution2(entries))
}
