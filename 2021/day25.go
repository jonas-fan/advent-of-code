package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func read(reader io.Reader) <-chan string {
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

type Position struct {
	Row int
	Col int
}

func solution(row int, col int, east map[Position]bool, south map[Position]bool) int {
	var next map[Position]bool
	var round int

	move := true

	for round = 0; move; round++ {
		move = false
		next = map[Position]bool{}

		for pos := range east {
			target := Position{
				Row: pos.Row,
				Col: (pos.Col + 1) % col,
			}

			if _, have := east[target]; have {
				next[pos] = true
			} else if _, have := south[target]; have {
				next[pos] = true
			} else {
				next[target] = true
				move = true
			}
		}

		east = next
		next = map[Position]bool{}

		for pos := range south {
			target := Position{
				Row: (pos.Row + 1) % row,
				Col: pos.Col,
			}

			if _, have := east[target]; have {
				next[pos] = true
			} else if _, have := south[target]; have {
				next[pos] = true
			} else {
				next[target] = true
				move = true
			}
		}

		south = next
	}

	return round
}

func main() {
	inputs := []string{}

	for input := range read(os.Stdin) {
		inputs = append(inputs, input)
	}

	east := map[Position]bool{}
	south := map[Position]bool{}

	for row, input := range inputs {
		for col, val := range input {
			pos := Position{
				Row: row,
				Col: col,
			}

			switch val {
			case '>':
				east[pos] = true
			case 'v':
				south[pos] = true
			}
		}
	}

	fmt.Println(solution(len(inputs), len(inputs[0]), east, south))
}
