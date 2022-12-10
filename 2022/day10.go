package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}

func atoi(str string) int {
	out, _ := strconv.Atoi(str)

	return out
}

type Instruction struct {
	op  string
	val int
	run int
}

func sol1(instructions []*Instruction) int {
	out := 0
	cycle := 0
	x := 1

	for i := 0; i < len(instructions); i++ {
		cycle++
		instructions[i].run++

		if (cycle >= 20) && ((cycle-20)%40 == 0) {
			out += cycle * x
		}

		switch instructions[i].op {
		case "noop":
			// do nothing
		case "addx":
			// 2 cycles required
			if instructions[i].run < 2 {
				i--
			} else {
				x += instructions[i].val
			}
		}
	}

	return out
}

func sol2(instructions []*Instruction) {
	crt := [][40]int{}

	for i := 0; i < 6; i++ {
		crt = append(crt, [40]int{})
	}

	cycle := 0
	x := 1

	for i := 0; i < len(instructions); i++ {
		cycle++
		instructions[i].run++

		pos := (cycle - 1) % 40

		if abs(pos-x) < 2 {
			crt[(cycle-1)/40][pos]++
		}

		switch instructions[i].op {
		case "noop":
			// do nothing
		case "addx":
			// 2 cycles required
			if instructions[i].run < 2 {
				i--
			} else {
				x += instructions[i].val
			}
		}
	}

	for row := 0; row < len(crt); row++ {
		for col := 0; col < len(crt[row]); col++ {
			if crt[row][col] > 0 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}

		fmt.Println()
	}
}

func parseInstructions(lines []string) []*Instruction {
	instructions := make([]*Instruction, 0, len(lines))

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		instruction := &Instruction{
			op: tokens[0],
		}

		if len(tokens) > 1 {
			instruction.val = atoi(tokens[1])
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}

func main() {
	lines := []string{}

	for input := range read(os.Stdin) {
		lines = append(lines, input)
	}

	fmt.Println("Ans1", sol1(parseInstructions(lines)))
	fmt.Println("Ans2")
	sol2(parseInstructions(lines))
}
