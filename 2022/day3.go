package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
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

func min(lhs int, rhs int) int {
	if lhs < rhs {
		return lhs
	}

	return rhs
}

func interacts(lhs string, rhs string) string {
	lseen := [256]int{}
	rseen := [256]int{}

	for _, char := range lhs {
		lseen[char]++
	}

	for _, char := range rhs {
		rseen[char]++
	}

	out := []rune{}

	for i := 0; i < len(lseen); i++ {
		if lseen[i] == 0 || rseen[i] == 0 {
			continue
		}

		out = append(out, rune(i))
	}

	return string(out)
}

func solution1(rucksacks []string) int {
	out := 0

	for _, rucksack := range rucksacks {
		mid := len(rucksack) >> 1
		lhs, rhs := rucksack[:mid], rucksack[mid:]

		for _, item := range interacts(lhs, rhs) {
			if unicode.IsLower(item) {
				out += int(item-'a') + 1
			} else if unicode.IsUpper(item) {
				out += int(item-'A') + 27
			}
		}
	}

	return out
}

func solution2(rucksacks []string) int {
	out := 0

	for i := 0; i < len(rucksacks); i += 3 {
		lhs := interacts(rucksacks[i], rucksacks[i+1])
		rhs := interacts(rucksacks[i+1], rucksacks[i+2])

		for _, item := range interacts(lhs, rhs) {
			if unicode.IsLower(item) {
				out += int(item-'a') + 1
			} else if unicode.IsUpper(item) {
				out += int(item-'A') + 27
			}
		}
	}

	return out
}

func main() {
	rucksacks := make([]string, 0)

	for input := range read(os.Stdin) {
		rucksacks = append(rucksacks, input)
	}

	fmt.Println(solution1(rucksacks))
	fmt.Println(solution2(rucksacks))
}
