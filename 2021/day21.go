package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type DeterministicDice struct {
	side    int
	current int
	rolls   int
}

func (d *DeterministicDice) Roll() int {
	d.rolls++
	d.current++

	out := d.current

	d.current %= d.side

	return out
}

func (d *DeterministicDice) Count() int {
	return d.rolls
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

func str2int(str string) int {
	val, _ := strconv.Atoi(str)

	return val
}

func max(lhs int, rhs int) int {
	if lhs < rhs {
		return rhs
	}

	return lhs
}

func sliceMax(nums []int) int {
	out := 0

	for _, each := range nums {
		out = max(out, each)
	}

	return out
}

func solution1(players []int) int {
	positions := make([]int, len(players))

	for i, pos := range players {
		positions[i] = pos
	}

	score := make([]int, len(players))
	d := &DeterministicDice{
		side: 100,
	}

Loop:
	for {
		for i := range positions {
			positions[i] += d.Roll() + d.Roll() + d.Roll()
			positions[i] %= 10

			if positions[i] > 0 {
				score[i] += positions[i]
			} else {
				score[i] += 10
			}

			if score[i] >= 1000 {
				break Loop
			}
		}
	}

	out := 0

	for _, each := range score {
		if each < 1000 {
			out += d.Count() * each
			break
		}
	}

	return out
}

func newSlice(row int, col int) [][]int {
	out := make([][]int, row)

	for i := range out {
		out[i] = make([]int, col)
	}

	return out
}

var steps = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func solution2(players []int) int {
	universes := 1
	win := make([]int, len(players))
	scores := [][][]int{newSlice(21, 10), newSlice(21, 10)}

	for i, start := range players {
		scores[i][0][start-1] = 1
	}

	for turn := 0; universes > 0; turn = (turn + 1) % len(scores) {
		ongoing := 0
		tmp := newSlice(21, 10)

		for score, positions := range scores[turn] {
			for pos, count := range positions {
				if count == 0 {
					continue
				}

				for step, times := range steps {
					next := (pos + step) % 10

					if score+next+1 >= len(scores[turn]) {
						win[turn] += count * times * universes
					} else {
						tmp[score+next+1][next] += count * times
						ongoing += count * times
					}
				}
			}
		}

		scores[turn] = tmp
		universes = ongoing
	}

	return sliceMax(win)
}

func main() {
	players := []int{}

	for input := range read(os.Stdin) {
		tokens := strings.Split(input, " ")

		players = append(players, str2int(tokens[len(tokens)-1]))

	}

	fmt.Println(solution1(players))
	fmt.Println(solution2(players))
}
