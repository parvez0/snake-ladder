package board

import "fmt"

type Snake map[int]int
type Ladder map[int]int

type Dimensions struct {
	M int
	N int
}

func (d *Dimensions) getBoardLen() int {
	return d.M * d.N
}

type Board interface {
	NextPosition(lastPosition, diceResult int) int
	BoardSize() int
	Snakes() Snake
	Ladders() Ladder
}

func NewBoard(args ...interface{}) Board {
	var (
		snakes = Snake{30:14, 41:20, 59:37, 67:50, 83:80, 92:70, 95:44, 99:5}
		ladders = Ladder{2:23, 8:12, 17:93, 29:54, 32:51, 39:80, 62:78, 75:96}
		dimensions = Dimensions{10,10}
	)
	for _, v := range args {
		switch v.(type) {
		case Dimensions:
			dimensions = v.(Dimensions)
		case Snake:
			snakes = v.(Snake)
		case Ladder:
			ladders = v.(Ladder)
		}
	}
	return board{
		Snake:     snakes,
		Ladder:    ladders,
		Size:      dimensions.getBoardLen(),
	}
}

type board struct {
	Snake      Snake
	Ladder     Ladder
	Size       int
}

func (b board) NextPosition(lastPos, diceRes int) int {
	newPos := lastPos + diceRes
	if v, e := b.Snake[newPos]; e {
		fmt.Printf("     Bit by snake %d -> %d\n", newPos, v)
		newPos = v
	}
	if v, e := b.Ladder[newPos]; e {
		fmt.Printf("     Climbed ladder %d -> %d\n", newPos, v)
		newPos = v
	}
	if newPos > b.Size {
		return lastPos
	}
	if newPos == b.Size {
		fmt.Printf("     Goal Reached\n")g
		return -1
	}
	return newPos
}

func (b board) BoardSize() int {
	return b.Size
}

func (b board) Snakes() Snake {
	return b.Snake
}

func (b board) Ladders() Ladder {
	return b.Ladder
}

