package main

import (
	"fmt"
	"os"

	"github.com/jonas-fan/advent-of-code/funcs"
)

func i2s(num int) string {
	digits := []int{}

	for num > 0 {
		digits = append(digits, num%5)
		num /= 5
	}

	for i := 0; i < len(digits); i++ {
		carry := 0

		if digits[i] == 3 {
			digits[i] = -2
			carry = 1
		} else if digits[i] == 4 {
			digits[i] = -1
			carry = 1
		} else if digits[i] > 4 {
			digits[i] %= 5
			carry = 1
		}

		if carry > 0 {
			if i+1 < len(digits) {
				digits[i+1] += carry
			} else {
				digits = append(digits, carry)
			}
		}
	}

	out := []rune{}

	for i := len(digits) - 1; i >= 0; i-- {
		switch digits[i] {
		case -2:
			out = append(out, '=')
		case -1:
			out = append(out, '-')
		default:
			out = append(out, rune(digits[i]+'0'))
		}
	}

	return string(out)
}

func s2i(num string) int {
	out := 0
	step := 1

	for i := len(num) - 1; i >= 0; i-- {
		switch num[i] {
		case '-':
			out += step * -1
		case '=':
			out += step * -2
		default:
			out += step * int(num[i]-'0')
		}

		step *= 5
	}

	return out
}

func sol1(lines []string) string {
	out := 0

	for _, line := range lines {
		out += s2i(line)
	}

	return i2s(out)
}

func main() {
	lines := []string{}

	for input := range funcs.ReadLines(os.Stdin) {
		lines = append(lines, input)
	}

	fmt.Println("Ans1", sol1(lines))
}
