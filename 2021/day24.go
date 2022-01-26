package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func min(lhs int, rhs int) int {
	if lhs < rhs {
		return lhs
	}

	return rhs
}

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}

	return lhs
}

func str2int(str string) int {
	val, _ := strconv.Atoi(str)

	return val
}

type Program struct {
	Iteration int
	Offsets   [][]int
	Divider   []int
}

func (p *Program) Compile(instructions []string) {
	for _, instruction := range instructions {
		if strings.HasPrefix(instruction, "inp") {
			p.Iteration++
		}
	}

	step := len(instructions) / p.Iteration

	for i := 0; i < len(instructions); i += step {
		offset := []int{
			str2int(strings.SplitN(instructions[i+5], " ", 3)[2]),
			str2int(strings.SplitN(instructions[i+15], " ", 3)[2]),
		}

		p.Offsets = append(p.Offsets, offset)
		p.Divider = append(p.Divider, str2int(strings.SplitN(instructions[i+4], " ", 3)[2]))
	}
}

func solution1(instructions []string) int {
	prog := &Program{}
	prog.Compile(instructions)

	result := map[int]int{
		0: 0,
	}

	for i := 0; i < prog.Iteration; i++ {
		fmt.Println(i, "digit is in progress")

		cache := map[int]int{}

		for z, model := range result {
			for w := 1; w < 10; w++ {
				num := model*10 + w
				val := z / prog.Divider[i]

				if ((z % 26) + prog.Offsets[i][0]) != w {
					val = val*26 + w + prog.Offsets[i][1]
				}

				if origin, have := cache[val]; have {
					cache[val] = max(origin, num)
				} else {
					cache[val] = num
				}
			}
		}

		result = cache
	}

	return result[0]
}

func solution2(instructions []string) int {
	prog := &Program{}
	prog.Compile(instructions)

	result := map[int]int{
		0: 0,
	}

	for i := 0; i < prog.Iteration; i++ {
		fmt.Println(i, "digit is in progress")

		cache := map[int]int{}

		for z, model := range result {
			for w := 1; w < 10; w++ {
				num := model*10 + w
				val := z / prog.Divider[i]

				if ((z % 26) + prog.Offsets[i][0]) != w {
					val = val*26 + w + prog.Offsets[i][1]
				}

				if origin, have := cache[val]; have {
					cache[val] = min(origin, num)
				} else {
					cache[val] = num
				}
			}
		}

		result = cache
	}

	return result[0]
}

func main() {
	instructions := []string{}

	for input := range read(os.Stdin) {
		instructions = append(instructions, input)
	}

	fmt.Println(solution1(instructions))
	fmt.Println(solution2(instructions))
}
