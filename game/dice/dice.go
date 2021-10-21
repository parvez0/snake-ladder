package dice

import (
	"math/rand"
	"time"
)

type NumberOfDice int

type Dice interface {
	Roll() int
	generateRandom(ch chan int)
}

func NewDice(args ...int) Dice {
	var nDice NumberOfDice = 1
	if len(args) > 0 {
		nDice = NumberOfDice(args[0])
	}
	return dice{DiceCount: nDice}
}

type dice struct {
	DiceCount NumberOfDice
}

func (d dice) generateRandom(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	ch <- rand.Intn(6) + 1
}

func (d dice) Roll() int {
	dc := int(d.DiceCount)
	ch :=  make(chan int, dc)
	for i:=0; i<dc; i++ {
		go d.generateRandom(ch)
	}
	steps := 0
	for i:=0; i<dc; i++ {
		select {
		case x := <- ch:
			steps += x
		}
	}
	return steps
}

