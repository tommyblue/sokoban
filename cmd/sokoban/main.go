package main

import (
	"github.com/tommyblue/sokoban/engine"
	"github.com/tommyblue/sokoban/game"
)

func main() {
	ge := game.InitGame()
	engine.MainLoop(ge)
}
