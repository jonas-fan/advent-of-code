package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func solution1(rounds [][2]byte) int {
	out := 0

	for _, shape := range rounds {
		lval := int(shape[0] - 'X')
		rval := int(shape[1] - 'X')

		if (lval+1)%3 == rval {
			// win
			out += rval + 1 + 6
		} else if lval == rval {
			// draw
			out += rval + 1 + 3
		} else if (rval+1)%3 == lval {
			// lose
			out += rval + 1
		}
	}

	return out
}

func solution2(rounds [][2]byte) int {
	out := 0

	for _, shape := range rounds {
		lval := int(shape[0] - 'X')

		switch shape[1] {
		case 'X':
			// lose
			out += (lval-1+3)%3 + 1
		case 'Y':
			// draw
			out += lval + 1 + 3
		case 'Z':
			// win
			out += (lval+1)%3 + 1 + 6
		}
	}

	return out
}

func main() {
	rounds := [][2]byte{}

	for input := range read(os.Stdin) {
		shape := [2]byte{input[0] - 'A' + 'X', input[2]}

		rounds = append(rounds, shape)
	}

	fmt.Println(solution1(rounds))
	fmt.Println(solution2(rounds))
}
