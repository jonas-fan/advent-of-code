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
			out <- strings.TrimSpace(scanner.Text())
		}
	}()

	return out
}

func str2int(str string) int {
	val, _ := strconv.Atoi(str)

	return val
}

func sum(nums []int) int {
	out := 0

	for _, each := range nums {
		out += each
	}

	return out
}

func solution(fish []int, days int) int {
	if days < 1 {
		return 0
	}

	daysToSpwan := make([]int, 9)

	for _, each := range fish {
		daysToSpwan[each]++
	}

	for days > 0 {
		spawn := 0
		shift := 0

		for ; shift < len(daysToSpwan); shift++ {
			if spawn = daysToSpwan[shift]; spawn > 0 {
				break
			}
		}

		if shift += 1; shift > days {
			break
		}

		for i := 0; i < len(daysToSpwan)-shift; i++ {
			daysToSpwan[i] = daysToSpwan[i+shift]
			daysToSpwan[i+shift] = 0
		}

		daysToSpwan[6] += spawn
		daysToSpwan[8] = spawn
		days -= shift
	}

	return sum(daysToSpwan)
}

func main() {
	fish := []int{}

	for input := range read(os.Stdin) {
		for _, each := range strings.Split(input, ",") {
			fish = append(fish, str2int(each))
		}
	}

	fmt.Println(solution(fish, 80))
	fmt.Println(solution(fish, 256))
}
