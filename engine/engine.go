package engine

import (
	"github.com/tommyblue/sokoban/game"
	"github.com/tommyblue/sokoban/ui"
)

func MainLoop(ge *game.GameEngine) {
	gui := ge.GUI
	if gui == nil {
		gui = &ui.GUI{}
	}
	gui.Init()
	defer gui.Close()

	ge.IsRunning = true
	for ge.IsRunning {
		ge.ManageInput()
		gui.Draw(ge.Game.CurrentLevel)
	}
}

// func updateStatus() {}
