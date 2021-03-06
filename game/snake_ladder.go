package game

import (
	"fmt"
	"github.com/parvez0/snake-ladder/game/board"
	"github.com/parvez0/snake-ladder/game/dice"
	"github.com/parvez0/snake-ladder/game/player"
	"github.com/parvez0/snake-ladder/queue"
)

type Game interface {
	Play()
}

func NewGame(players, numDice int, args ...interface{}) Game {
	b := board.NewBoard(args...)
	d := dice.NewDice(numDice)
	q := queue.NewQueue()
	for i:=1; i<=players; i++ {
		p := player.NewPlayer(fmt.Sprintf("Player-%d", i))
		q.Push(p)
	}
	return snakeLadder{
		Board:   b,
		Dice:    d,
		Players: q,
	}
}

type snakeLadder struct {
	Board    board.Board
	Dice     dice.Dice
	Players  queue.Q
}

func (s snakeLadder) Play() {
	for s.Players.Len() > 0 {
		top := s.Players.Pop().(player.Player)
		fmt.Printf("%s is playing\n", top.GetName())
		total := s.Dice.Roll()
		position := s.Board.NextPosition(top.CurrentPosition(), total)
		if position == -1 {
			fmt.Printf("**** %s wins after %d moves ****\n", top.GetName(), top.MovesPlayed() + 1)
			continue
		}
		top.NewPosition(position)
		s.Players.Push(top)
	}
}
