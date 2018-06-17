package main

import (
	"github.com/tommyblue/sokoban/game"
)

func main() {
	ge := game.InitGame()
	defer ge.GUI.Close()
	game.MainLoop(ge)
}
