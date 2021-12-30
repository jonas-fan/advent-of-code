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

func enhance(image [][]int, algo []int, padding int) ([][]int, int) {
	out := make([][]int, len(image)+2)

	for i := range out {
		out[i] = make([]int, len(image)+2)
	}

	for row := range out {
		for col := range out[row] {
			value := 0

			for i := row - 2; i < row+1; i++ {
				for j := col - 2; j < col+1; j++ {
					bit := 0

					switch {
					case i < 0 || i >= len(image):
						bit = padding
					case j < 0 || j >= len(image[i]):
						bit = padding
					default:
						bit = image[i][j]
					}

					value = (value << 1) | bit
				}
			}

			out[row][col] = algo[value]
		}
	}

	if padding == 0 {
		padding = algo[padding]
	} else {
		padding = algo[len(algo)-1]
	}

	return out, padding
}

func lit(value int) bool {
	return value > 0
}

func count(list []int, comp func(value int) bool) int {
	out := 0

	for _, each := range list {
		if comp(each) {
			out++
		}
	}

	return out
}

func solution(image [][]int, algo []int, times int) int {
	padding := 0

	for n := 0; n < times; n++ {
		image, padding = enhance(image, algo, padding)
	}

	lights := 0

	for _, line := range image {
		lights += count(line, lit)
	}

	return lights
}

func parse(token string) []int {
	out := []int{}

	for _, each := range token {
		if each == '#' {
			out = append(out, 1)
		} else {
			out = append(out, 0)
		}
	}

	return out
}

func main() {
	algo := []int{}
	image := [][]int{}
	reader := read(os.Stdin)

	for input := range reader {
		if input == "" {
			break
		}

		algo = parse(input)
	}

	for input := range reader {
		image = append(image, parse(input))
	}

	fmt.Println(solution(image, algo, 2))
	fmt.Println(solution(image, algo, 50))
}
