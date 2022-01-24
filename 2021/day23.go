package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
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

var (
	Hallway = []Position{
		Position{X: 1, Y: 1},
		Position{X: 2, Y: 1},
		Position{X: 4, Y: 1},
		Position{X: 6, Y: 1},
		Position{X: 8, Y: 1},
		Position{X: 10, Y: 1},
		Position{X: 11, Y: 1},
	}

	Rooms = [][]Position{
		[]Position{Position{X: 3, Y: 2}, Position{X: 3, Y: 3}, Position{X: 3, Y: 4}, Position{X: 3, Y: 5}},
		[]Position{Position{X: 5, Y: 2}, Position{X: 5, Y: 3}, Position{X: 5, Y: 4}, Position{X: 5, Y: 5}},
		[]Position{Position{X: 7, Y: 2}, Position{X: 7, Y: 3}, Position{X: 7, Y: 4}, Position{X: 7, Y: 5}},
		[]Position{Position{X: 9, Y: 2}, Position{X: 9, Y: 3}, Position{X: 9, Y: 4}, Position{X: 9, Y: 5}},
	}
)

type Amphiod byte

func (a *Amphiod) Type() int {
	return int(*a - 'A')
}

func (a *Amphiod) Cost() int {
	switch kind := a.Type(); kind {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		panic(fmt.Sprintf("Invalid Amphiod: %q (type: %d)", *a, kind))
	}
}

type Position struct {
	X int
	Y int
}

func (p *Position) Offset() int {
	return p.Y*14 + p.X
}

func (p *Position) Distance(other Position) int {
	return abs(other.X-p.X) + abs(other.Y-p.Y)
}

type Burrow struct {
	Hallway map[Position]Amphiod
	Room    map[Position]Amphiod
	Depth   int
}

func (b *Burrow) Top(room int) (Amphiod, Position, bool) {
	r := Rooms[room]

	for i := 0; i < b.Depth; i++ {
		pos := r[i]

		if amphiod, ok := b.Room[pos]; ok {
			return amphiod, pos, true
		}
	}

	return 0, Position{X: 0, Y: 0}, false
}

func (b *Burrow) Settled(room int) bool {
	r := Rooms[room]

	for i := 0; i < b.Depth; i++ {
		pos := r[i]

		if amphiod, ok := b.Room[pos]; ok {
			if amphiod.Type() != room {
				return false
			}
		}
	}

	return true
}

func (b *Burrow) Available(room int) (Position, bool) {
	r := Rooms[room]

	for i := b.Depth - 1; i >= 0; i-- {
		pos := r[i]

		if _, have := b.Room[pos]; !have {
			return pos, true
		}
	}

	return Position{X: 0, Y: 0}, false
}

func (b *Burrow) Reachable(from Position, to Position) bool {
	var have bool

	if _, have = b.Hallway[to]; have {
		return false
	} else if _, have = b.Room[to]; have {
		return false
	}

	lower := min(from.X, to.X)
	upper := max(from.X, to.X)

	for _, pos := range Hallway {
		if pos == from {
			continue
		} else if pos.X <= lower {
			continue
		} else if pos.X > upper {
			break
		} else if _, have = b.Hallway[pos]; have {
			return false
		}
	}

	return true
}

func (b *Burrow) Move() map[*Burrow]int {
	moves := map[*Burrow]int{}

	for index := range Rooms {
		if b.Settled(index) {
			continue
		}

		amphiod, pos, ok := b.Top(index)

		if !ok {
			continue
		}

		for _, to := range Hallway {
			if b.Reachable(pos, to) {
				next := b.Copy()

				next.Hallway[to] = amphiod
				delete(next.Room, pos)

				moves[next] = pos.Distance(to) * amphiod.Cost()
			}
		}
	}

	for _, pos := range Hallway {
		if amphiod, have := b.Hallway[pos]; have {
			kind := amphiod.Type()

			if !b.Settled(kind) {
				continue
			}

			if to, ok := b.Available(kind); ok {
				if b.Reachable(pos, to) {
					next := b.Copy()

					next.Room[to] = amphiod
					delete(next.Hallway, pos)

					moves[next] = pos.Distance(to) * amphiod.Cost()
				}
			}
		}
	}

	return moves
}

func (b *Burrow) Copy() *Burrow {
	other := NewBurrow(b.Depth)

	for pos, amphiod := range b.Hallway {
		other.Hallway[pos] = amphiod
	}

	for pos, amphiod := range b.Room {
		other.Room[pos] = amphiod
	}

	return other
}

func (b *Burrow) String() string {
	var graph []byte

	switch b.Depth {
	case 2:
		graph = []byte(`#############
#...........#
###.#.#.#.###
  #.#.#.#.#  
  #########  `)
	case 4:
		graph = []byte(`#############
#...........#
###.#.#.#.###
  #.#.#.#.#  
  #.#.#.#.#  
  #.#.#.#.#  
  #########  `)
	}

	for pos, amphiod := range b.Hallway {
		graph[pos.Offset()] = byte(amphiod)
	}

	for pos, amphiod := range b.Room {
		graph[pos.Offset()] = byte(amphiod)
	}

	return string(graph)
}

func NewBurrow(depth int) *Burrow {
	return &Burrow{
		Hallway: make(map[Position]Amphiod),
		Room:    make(map[Position]Amphiod),
		Depth:   depth,
	}
}

func solution(graph string, goal string, depth int) int {
	burrow := NewBurrow(depth)

	for _, room := range Rooms {
		for i := 0; i < burrow.Depth; i++ {
			pos := room[i]

			burrow.Room[pos] = Amphiod(graph[pos.Offset()])
		}
	}

	cost := map[string]int{burrow.String(): 0}
	stack := []*Burrow{burrow}

	for len(stack) > 0 {
		burrow, stack = stack[len(stack)-1], stack[:len(stack)-1]
		baseCost := cost[burrow.String()]

		for next, energy := range burrow.Move() {
			graph = next.String()
			energy += baseCost

			if val, ok := cost[graph]; ok {
				if energy < val {
					cost[graph] = energy
					stack = append(stack, next)
				}
			} else {
				cost[graph] = energy
				stack = append(stack, next)
			}
		}
	}

	return cost[goal]
}

func main() {
	var graph string
	var goal string

	for input := range read(os.Stdin) {
		graph += input + "\n"
	}

	goal = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#  
  #########  `

	fmt.Println(solution(graph, goal, 2))

	graph = graph[:41] + `
  #D#C#B#A#  
  #D#B#A#C#  ` + graph[41:]

	goal = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#  
  #A#B#C#D#  
  #A#B#C#D#  
  #########  `

	fmt.Println(solution(graph, goal, 4))
}
