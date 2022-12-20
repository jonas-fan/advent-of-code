package main

import (
	"fmt"
	"os"

	"github.com/jonas-fan/advent-of-code/funcs"
)

type Node struct {
	prev *Node
	next *Node
	val  int
}

func sol(nums []int, mul int, rounds int) int {
	var order = []*Node{}
	var list *Node

	for i := 0; i < len(nums); i++ {
		node := &Node{
			val: nums[i] * mul,
		}

		order = append(order, node)
	}

	for i := 0; i < len(order); i++ {
		order[i].prev = order[(i-1+len(order))%len(order)]
		order[i].next = order[(i+1)%len(order)]

		if order[i].val == 0 {
			list = order[i]
		}
	}

	for r := 0; r < rounds; r++ {
		for i := 0; i < len(order); i++ {
			node := order[i]
			step := node.val % (len(order) - 1)

			if step < 0 {
				step += len(order) - 1
			}

			for ; step > 0; step-- {
				next := node.next

				node.prev.next = node.next
				node.next.prev = node.prev
				node.next = next.next
				node.next.prev = node
				node.prev = next
				node.prev.next = node
			}
		}
	}

	out := 0
	node := list

	for i := 1; i <= 3000; i++ {
		node = node.next

		if i%1000 == 0 {
			out += node.val
		}
	}

	return out
}

func main() {
	lines := []string{}

	for input := range funcs.ReadLines(os.Stdin) {
		lines = append(lines, input)
	}

	nums := funcs.Atois(lines)

	fmt.Println("Ans1", sol(nums, 1, 1))
	fmt.Println("Ans2", sol(nums, 811589153, 10))
}
