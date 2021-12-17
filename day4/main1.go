package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	DirectUp = 0x1 << iota
	DirectDown
	DirectLeft
	DirectRight
)

type Point struct {
	x int
	y int
}

type Board struct {
	board    [][]bool
	position map[int]Point
}

func (b *Board) bingo(num int) bool {
	if point, ok := b.position[num]; ok {
		b.board[point.y][point.x] = true

		switch 4 {
		case b.count(point, DirectUp) + b.count(point, DirectDown):
			return true
		case b.count(point, DirectLeft) + b.count(point, DirectRight):
			return true
			// case b.count(point, DirectUp | DirectRight) + b.count(point, DirectDown | DirectLeft):
			// 	return true
			// case b.count(point, DirectUp | DirectLeft) + b.count(point, DirectDown | DirectRight):
			// 	return true
		}
	}

	return false
}

func (b *Board) count(pos Point, direct int) int {
	switch {
	case direct&DirectUp > 0:
		pos.y--
	case direct&DirectDown > 0:
		pos.y++
	}

	switch {
	case direct&DirectLeft > 0:
		pos.x--
	case direct&DirectRight > 0:
		pos.x++
	}

	if pos.y < 0 || pos.y >= len(b.board) {
		return 0
	} else if pos.x < 0 || pos.x >= len(b.board[0]) {
		return 0
	} else if !b.board[pos.y][pos.x] {
		return 0
	}

	return 1 + b.count(pos, direct)
}

func (b *Board) unmarked() int {
	sum := 0

	for num, point := range b.position {
		if !b.board[point.y][point.x] {
			sum += num
		}
	}

	return sum
}

func FromTable(table [][]int) *Board {
	board := make([][]bool, len(table))
	position := make(map[int]Point)

	for row, vals := range table {
		board[row] = make([]bool, len(vals))

		for col, val := range vals {
			position[val] = Point{
				x: col,
				y: row,
			}
		}
	}

	return &Board{
		board:    board,
		position: position,
	}
}

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

func str2int(strs []string) []int {
	out := make([]int, 0, len(strs))

	for _, each := range strs {
		val, _ := strconv.Atoi(each)

		out = append(out, val)
	}

	return out
}

func solution(boards []*Board, steps []int) int {
	for _, step := range steps {
		for _, board := range boards {
			if board.bingo(step) {
				return board.unmarked() * step
			}
		}
	}

	return 0
}

func main() {
	steps := []int{}
	table := [][]int{}
	boards := []*Board{}

	for input := range read(os.Stdin) {
		if len(steps) == 0 {
			steps = str2int(strings.Split(input, ","))
		} else if input != "" {
			table = append(table, str2int(strings.Fields(input)))
		} else if len(table) > 0 {
			boards = append(boards, FromTable(table))
			table = [][]int{}
		}
	}

	fmt.Println(solution(boards, steps))
}
