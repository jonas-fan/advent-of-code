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

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}

	return lhs
}

type Pair struct {
	Left  interface{}
	Right interface{}
}

type Number struct {
	Value int
}

func ladd(node interface{}, value int) {
	switch node.(type) {
	case *Pair:
		ladd(node.(*Pair).Left, value)
	case *Number:
		node.(*Number).Value += value
	}
}

func radd(node interface{}, value int) {
	switch node.(type) {
	case *Pair:
		radd(node.(*Pair).Right, value)
	case *Number:
		node.(*Number).Value += value
	}
}

func explode(pair *Pair, depth int) (bool, int, int) {
	if lnum, ok := pair.Left.(*Number); !ok {
		// not a leaf
	} else if rnum, ok := pair.Right.(*Number); !ok {
		// not a leaf
	} else {
		return false, lnum.Value, rnum.Value
	}

	if lpair, ok := pair.Left.(*Pair); ok {
		exploded, llvalue, lrvalue := explode(lpair, depth+1)

		if !exploded && depth > 2 {
			pair.Left = &Number{Value: 0}
			exploded = true
		}

		if exploded {
			ladd(pair.Right, lrvalue)

			return exploded, llvalue, 0
		}
	}

	if rpair, ok := pair.Right.(*Pair); ok {
		exploded, rlvalue, rrvalue := explode(rpair, depth+1)

		if !exploded && depth > 2 {
			pair.Right = &Number{Value: 0}
			exploded = true
		}

		if exploded {
			radd(pair.Left, rlvalue)

			return exploded, 0, rrvalue
		}
	}

	return false, 0, 0
}

func split(pair *Pair) bool {
	switch pair.Left.(type) {
	case *Pair:
		if split(pair.Left.(*Pair)) {
			return true
		}
	case *Number:
		if lvalue := pair.Left.(*Number).Value; lvalue >= 10 {
			pair.Left = &Pair{
				Left:  &Number{Value: lvalue >> 1},
				Right: &Number{Value: lvalue - (lvalue >> 1)},
			}

			return true
		}
	}

	switch pair.Right.(type) {
	case *Pair:
		if split(pair.Right.(*Pair)) {
			return true
		}
	case *Number:
		if rvalue := pair.Right.(*Number).Value; rvalue >= 10 {
			pair.Right = &Pair{
				Left:  &Number{Value: rvalue >> 1},
				Right: &Number{Value: rvalue - (rvalue >> 1)},
			}

			return true
		}
	}

	return false
}

func reduce(pair *Pair) {
	for {
		for {
			updated, _, _ := explode(pair, 0)

			if !updated {
				break
			}
		}

		if !split(pair) {
			break
		}
	}
}

func value(node interface{}) int {
	if number, ok := node.(*Number); ok {
		return number.Value
	}

	return 3*value(node.(*Pair).Left) + 2*value(node.(*Pair).Right)
}

func parse(line string) *Pair {
	nums := []interface{}{}
	brackets := []rune{}

	for _, each := range line {
		switch each {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			nums = append(nums, &Number{Value: int(each - '0')})
		case '[':
			brackets = append(brackets, each)
		case ']':
			pair := nums[len(nums)-2:]

			nums = nums[:len(nums)-2]
			nums = append(nums, &Pair{
				Left:  pair[0],
				Right: pair[1],
			})

			brackets = brackets[:len(brackets)-1]
		}
	}

	return nums[0].(*Pair)
}

func solution1(lines []string) int {
	var head *Pair

	for _, line := range lines {
		if head != nil {
			head = &Pair{
				Left:  head,
				Right: parse(line),
			}
		} else {
			head = parse(line)
		}

		reduce(head)
	}

	return value(head)
}

func solution2(lines []string) int {
	magnitude := 0

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines); j++ {
			if i == j {
				continue
			}

			pair := &Pair{
				Left:  parse(lines[i]),
				Right: parse(lines[j]),
			}

			reduce(pair)

			magnitude = max(magnitude, value(pair))
		}
	}

	return magnitude
}

func main() {
	lines := []string{}

	for input := range read(os.Stdin) {
		lines = append(lines, input)
	}

	fmt.Println(solution1(lines))
	fmt.Println(solution2(lines))
}
