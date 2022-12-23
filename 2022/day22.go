package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jonas-fan/advent-of-code/funcs"
)

const MAX_DIRECT = 4
const MAX_WIDTH = 50

// Direct:
// - 0: >
// - 1: v
// - 2: <
// - 3: ^
type Position struct {
	row    int
	col    int
	direct int
}

func (p *Position) Next() Position {
	next := *p

	switch p.direct {
	case 0:
		next.col++
	case 1:
		next.row++
	case 2:
		next.col--
	case 3:
		next.row--
	}

	return next
}

func sol1(m [][]rune, steps []string) int {
	pos := Position{}

	for col := 0; col < len(m[0]); col++ {
		if m[0][col] == '.' {
			pos.col = col
			break
		}
	}

	for _, step := range steps {
		switch step {
		case "L":
			pos.direct = (pos.direct - 1 + MAX_DIRECT) % MAX_DIRECT
		case "R":
			pos.direct = (pos.direct + 1 + MAX_DIRECT) % MAX_DIRECT
		default:
			last := pos
			moves := funcs.Atoi(step)

			for i := 0; i < moves; i++ {
				next := pos.Next()

				next.row = (next.row + len(m)) % len(m)
				next.col = (next.col + len(m[0])) % len(m[0])

				if m[next.row][next.col] == '#' {
					pos = last
					break
				} else if m[next.row][next.col] == '.' {
					pos = next
					last = pos
				} else {
					pos = next
					i--
				}
			}
		}
	}

	return 1000*(pos.row+1) + 4*(pos.col+1) + pos.direct
}

const (
	R int = iota
	B
	L
	T
)

type Face struct {
	id   int
	side int
}

// Example,
//
// var faceMap = map[int][4]Face{
// 	//         >           v           <           ^
// 	1: [4]Face{Face{6, R}, Face{4, T}, Face{3, T}, Face{2, T}},
// 	2: [4]Face{Face{3, L}, Face{5, B}, Face{6, B}, Face{1, T}},
// 	3: [4]Face{Face{4, L}, Face{5, L}, Face{2, R}, Face{1, L}},
// 	4: [4]Face{Face{6, T}, Face{5, T}, Face{3, R}, Face{1, B}},
// 	5: [4]Face{Face{6, L}, Face{2, B}, Face{3, B}, Face{4, B}},
// 	6: [4]Face{Face{1, R}, Face{2, L}, Face{5, R}, Face{4, R}},
// }

// var faces = [4][4]int{
// 	[4]int{0, 0, 1, 0},
// 	[4]int{2, 3, 4, 0},
// 	[4]int{0, 0, 5, 6},
// 	[4]int{0, 0, 0, 0},
// }

var faceMap = map[int][4]Face{
	//         >           v           <           ^
	1: [4]Face{Face{2, L}, Face{3, T}, Face{4, L}, Face{6, L}},
	2: [4]Face{Face{5, R}, Face{3, R}, Face{1, R}, Face{6, B}},
	3: [4]Face{Face{2, B}, Face{5, T}, Face{4, T}, Face{1, B}},
	4: [4]Face{Face{5, L}, Face{6, T}, Face{1, L}, Face{3, L}},
	5: [4]Face{Face{2, R}, Face{6, R}, Face{4, R}, Face{3, B}},
	6: [4]Face{Face{5, B}, Face{2, T}, Face{1, T}, Face{4, B}},
}

var faces = [4][4]int{
	[4]int{0, 1, 2, 0},
	[4]int{0, 3, 0, 0},
	[4]int{4, 5, 0, 0},
	[4]int{6, 0, 0, 0},
}

func posToface(pos Position) int {
	row, col := pos.row/MAX_WIDTH, pos.col/MAX_WIDTH

	if faces[row][col] > 0 {
		return faces[row][col]
	}

	panic("unknown position")
}

func faceToPos(face int) Position {
	for row := 0; row < len(faces); row++ {
		for col := 0; col < len(faces[row]); col++ {
			if faces[row][col] == face {
				return Position{
					row: row * MAX_WIDTH,
					col: col * MAX_WIDTH,
				}
			}
		}
	}

	panic("unknown face")
}

func sol2(m [][]rune, steps []string) int {
	pos := Position{}

	for col := 0; col < len(m[0]); col++ {
		if m[0][col] == '.' {
			pos.col = col
			break
		}
	}

	for _, step := range steps {
		switch step {
		case "L":
			pos.direct = (pos.direct - 1 + MAX_DIRECT) % MAX_DIRECT
		case "R":
			pos.direct = (pos.direct + 1 + MAX_DIRECT) % MAX_DIRECT
		default:
			last := pos
			moves := funcs.Atoi(step)

			for i := 0; i < moves; i++ {
				next := pos.Next()

				next.row = (next.row + len(m)) % len(m)
				next.col = (next.col + len(m[0])) % len(m[0])

				if m[next.row][next.col] == '#' {
					pos = last
					break
				} else if m[next.row][next.col] == '.' {
					pos = next
					last = pos
				} else {
					nextFace := faceMap[posToface(pos)][pos.direct]
					next = faceToPos(nextFace.id)

					switch nextFace.side {
					case R:
						next.direct = 2
						next.col += MAX_WIDTH
						switch pos.direct {
						case 0:
							next.row += MAX_WIDTH - pos.row%MAX_WIDTH - 1
						case 1:
							next.row += pos.col % MAX_WIDTH
						case 2:
							next.row += pos.row % MAX_WIDTH
						case 3:
							next.row += MAX_WIDTH - pos.col%MAX_WIDTH - 1
						}
					case B:
						next.direct = 3
						next.row += MAX_WIDTH
						switch pos.direct {
						case 0:
							next.col += pos.row % MAX_WIDTH
						case 1:
							next.col += MAX_WIDTH - pos.col%MAX_WIDTH - 1
						case 2:
							next.col += MAX_WIDTH - pos.row%MAX_WIDTH - 1
						case 3:
							next.col += pos.col % MAX_WIDTH
						}
					case L:
						next.direct = 0
						next.col--
						switch pos.direct {
						case 0:
							next.row += pos.row % MAX_WIDTH
						case 1:
							next.row += MAX_WIDTH - pos.col%MAX_WIDTH - 1
						case 2:
							next.row += MAX_WIDTH - pos.row%MAX_WIDTH - 1
						case 3:
							next.row += pos.col % MAX_WIDTH
						}
					case T:
						next.direct = 1
						next.row--
						switch pos.direct {
						case 0:
							next.col += MAX_WIDTH - pos.row%MAX_WIDTH - 1
						case 1:
							next.col += pos.col % MAX_WIDTH
						case 2:
							next.col += pos.row % MAX_WIDTH
						case 3:
							next.col += MAX_WIDTH - pos.col%MAX_WIDTH - 1
						}
					}

					pos = next
					i--
				}
			}
		}
	}

	return 1000*(pos.row+1) + 4*(pos.col+1) + pos.direct
}

func main() {
	m := [][]rune{}
	steps := []string{}

	{
		width := 0
		lines := []string{}
		in := funcs.ReadLines(os.Stdin)

		for input := range in {
			if input == "" {
				break
			}

			lines = append(lines, input)
			width = funcs.Max(width, len(input))
		}

		for i := 0; i < MAX_WIDTH*MAX_WIDTH; i++ {
			row := make([]rune, width)

			if i < len(lines) {
				copy(row, []rune(lines[i]))
			}

			m = append(m, row)
		}

		for input := range in {
			begin := 0

			for end := 0; end < len(input); end++ {
				if strings.ContainsAny(input[end:end+1], "LR") {
					steps = append(steps, input[begin:end])
					steps = append(steps, input[end:end+1])
					begin = end + 1
				}
			}

			if begin != len(input) {
				steps = append(steps, input[begin:])
			}
		}
	}

	fmt.Println("Ans1", sol1(m, steps))
	fmt.Println("Ans2", sol2(m, steps))
}
