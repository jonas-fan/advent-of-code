package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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

func atoi(str string) int {
	out, _ := strconv.Atoi(str)

	return out
}

type Tree struct {
	parent   *Tree
	children []*Tree
	val      int
}

func makeTree(node *Tree, lines []string) {
	if len(lines) == 0 {
		return
	}

	tokens := strings.Split(lines[0], " ")

	lines = lines[1:]

	if tokens[0] == "$" {
		if tokens[1] == "cd" {
			if tokens[2] == "." {
				// current path
			} else if tokens[2] == ".." {
				// go back
				node = node.parent
			} else {
				// go next
				child := &Tree{parent: node}

				node.children = append(node.children, child)
				node = child
			}
		} else if tokens[1] == "ls" {
			for len(lines) > 0 && lines[0][0] != '$' {
				node.val += atoi(strings.Split(lines[0], " ")[0])
				lines = lines[1:]
			}
		}
	}

	makeTree(node, lines)
}

func walk(node *Tree, out []int) []int {
	if node == nil {
		return out
	}

	sum := node.val

	for _, child := range node.children {
		out = walk(child, out)

		sum += out[len(out)-1]
	}

	return append(out, sum)
}

func sol1(lines []string) int {
	top := &Tree{}

	makeTree(top, lines)

	sizes := walk(top, []int{})

	sort.Ints(sizes)

	out := 0

	for _, size := range sizes {
		if size > 100000 {
			break
		}

		out += size
	}

	return out
}

func sol2(lines []string) int {
	top := &Tree{}

	makeTree(top, lines)

	sizes := walk(top, []int{})

	sort.Ints(sizes)

	sum := sizes[len(sizes)-1]

	for _, size := range sizes {
		if sum-size < 40000000 {
			return size
		}
	}

	return 0
}

func main() {
	lines := []string{}

	for input := range read(os.Stdin) {
		lines = append(lines, input)
	}

	fmt.Println("Ans1", sol1(lines))
	fmt.Println("Ans2", sol2(lines))
}
