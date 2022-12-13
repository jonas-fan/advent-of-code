package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
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

func atoi(str string) int {
	out, _ := strconv.Atoi(str)

	return out
}

type Packet interface{}

func findLast(packets []Packet, target string) int {
	i := len(packets) - 1

	for ; i >= 0; i-- {
		if val, ok := packets[i].(string); ok {
			if val == "[" {
				return i
			}
		}
	}

	return i
}

func parse(line string) Packet {
	stack := []Packet{}
	buffer := []rune{}

	for _, char := range line {
		switch char {
		case '[':
			// push onto stack
			stack = append(stack, string(char))

		case ']':
			// flush and push onto stack
			if len(buffer) > 0 {
				stack = append(stack, atoi(string(buffer)))
				buffer = buffer[:0]
			}

			// squash
			index := findLast(stack, "[")

			stack[index] = append([]Packet{}, stack[index+1:]...)
			stack = stack[:index+1]

		case ',':
			// flush and push onto stack
			if len(buffer) > 0 {
				stack = append(stack, atoi(string(buffer)))
				buffer = buffer[:0]
			}

		default:
			// push into buffer
			buffer = append(buffer, char)
		}
	}

	return stack[0]
}

func compare(lhs []Packet, rhs []Packet) int {
	for len(lhs) > 0 && len(rhs) > 0 {
		lval := lhs[0]
		rval := rhs[0]

		switch lval.(type) {
		case []Packet:
			switch rval.(type) {
			case []Packet:
				if ok := compare(lval.([]Packet), rval.([]Packet)); ok != 0 {
					return ok
				}

			case int:
				if ok := compare(lval.([]Packet), []Packet{rval}); ok != 0 {
					return ok
				}
			}

		case int:
			switch rval.(type) {
			case []Packet:
				if ok := compare([]Packet{lval}, rval.([]Packet)); ok != 0 {
					return ok
				}

			case int:
				if lval.(int) < rval.(int) {
					return -1
				} else if lval.(int) > rval.(int) {
					return 1
				}
			}
		}

		lhs = lhs[1:]
		rhs = rhs[1:]
	}

	if len(lhs) < len(rhs) {
		return -1
	} else if len(lhs) > len(rhs) {
		return 1
	}

	return 0
}

func sol1(lhs []string, rhs []string) int {
	out := 0

	for i := range lhs {
		lval := []Packet{parse(lhs[i])}
		rval := []Packet{parse(rhs[i])}

		if compare(lval, rval) < 0 {
			out += i + 1
		}
	}

	return out
}

func sol2(lhs []string, rhs []string) int {
	packets := []Packet{}

	for _, each := range append(lhs, rhs...) {
		packets = append(packets, parse(each))
	}

	extras := []Packet{
		parse("[[2]]"),
		parse("[[6]]"),
	}

	for _, each := range extras {
		packets = append(packets, each)
	}

	sort.Slice(packets, func(i int, j int) bool {
		lval := []Packet{packets[i]}
		rval := []Packet{packets[j]}

		return compare(lval, rval) < 0
	})

	out := 1

	for i, each := range packets {
		lval := []Packet{each}

		for _, extra := range extras {
			rval := []Packet{extra}

			if compare(lval, rval) == 0 {
				out *= i + 1
				break
			}
		}
	}

	return out
}

func main() {
	lines := []string{}

	for input := range read(os.Stdin) {
		if input != "" {
			lines = append(lines, input)
		}
	}

	lhs := []string{}
	rhs := []string{}

	for i := 0; i < len(lines); i += 2 {
		lhs = append(lhs, lines[i])
		rhs = append(rhs, lines[i+1])
	}

	fmt.Println("Ans1", sol1(lhs, rhs))
	fmt.Println("Ans2", sol2(lhs, rhs))
}
