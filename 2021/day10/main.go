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

var scopeMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func solution1(syntaxes [][]rune) int {
	var pointMap = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	points := 0

	for _, syntax := range syntaxes {
		stack := []rune{}

	Loop:
		for _, each := range syntax {
			switch each {
			case '(', '[', '{', '<':
				stack = append(stack, each)
			default:
				top := len(stack) - 1

				if top < 0 {
					break Loop
				} else if scopeMap[stack[top]] != each {
					points += pointMap[each]
					break Loop
				}

				stack = stack[:top]
			}
		}
	}

	return points
}

func solution2(syntaxes [][]rune) int {
	var pointMap = map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	points := []int{}

	for _, syntax := range syntaxes {
		valid := true
		stack := []rune{}

	Loop:
		for _, each := range syntax {
			switch each {
			case '(', '[', '{', '<':
				stack = append(stack, each)
			default:
				top := len(stack) - 1

				if top < 0 {
					valid = false
					break Loop
				} else if scopeMap[stack[top]] != each {
					valid = false
					break Loop
				}

				stack = stack[:top]
			}
		}

		if valid {
			point := 0

			for i := len(stack) - 1; i >= 0; i-- {
				point = point*5 + pointMap[stack[i]]
			}

			points = append(points, point)
		}
	}

	sort.Slice(points, func(lhs int, rhs int) bool {
		return points[lhs] < points[rhs]
	})

	return points[len(points)>>1]
}

func main() {
	syntaxes := [][]rune{}

	for input := range read(os.Stdin) {
		syntaxes = append(syntaxes, []rune(input))
	}

	fmt.Println(solution1(syntaxes))
	fmt.Println(solution2(syntaxes))
}
