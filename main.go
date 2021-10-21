package main

import (
	"github.com/parvez0/snake-ladder/game"
	"github.com/parvez0/snake-ladder/game/board"
)

func main() {
	dm := board.Dimensions{M: 10, N: 10}
	sn := board.Snake{30:14, 41:20, 59:37, 67:50, 83:80, 92:70, 95:44, 99:5}
	ld := board.Ladder{2:23, 8:12, 17:93, 29:54, 32:51, 39:80, 62:78, 75:96}
	g := game.NewGame(3, 3, dm, sn, ld)
	g.Play()
}
