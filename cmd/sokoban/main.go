package main

import (
	"github.com/tommyblue/sokoban/engine"
	"github.com/tommyblue/sokoban/game"
)

func main() {
	g := game.InitGame()
	engine.MainLoop(g)
}
