package dice

import "math/rand"

type NumberOfDice int

type Dice interface {
	Roll() int
}

func NewDice(args ...int) Dice {
	var nDice NumberOfDice = 1
	if len(args) > 0 {
		nDice = NumberOfDice(args[0])
	}
	return dice{NumDice: nDice}
}

type dice struct {
	NumDice NumberOfDice
}

func (d dice) Roll() int {
	lmt := int(6 * d.NumDice)
	return rand.Intn(lmt)
}

