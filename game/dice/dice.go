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
	return dice{NumDice: nDice}
}

type dice struct {
	NumDice NumberOfDice
}

func (d dice) generateRandom(ch chan int) {
	rand.Seed(time.Now().UnixNano())
	ch <- rand.Intn(6) + 1
}

func (d dice) Roll() int {
	ch :=  make(chan int, d.NumDice)
	for i:=0; i<int(d.NumDice); i++ {
		go d.generateRandom(ch)
	}
	steps := 0
	for i:=0; i<int(d.NumDice); i++ {
		select {
		case x := <- ch:
			steps += x
		}
	}
	return steps
}

