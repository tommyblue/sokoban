package game

import (
	"github.com/tommyblue/sokoban"
	"github.com/tommyblue/sokoban/ui"
)

// Engine represents the game
type Engine struct {
	GameState *sokoban.GameState
	Game      *sokoban.Game
	// TODO: Remove ui dependency (move to sokoban?)
	GUI *ui.GUI
}

// InitGame initializes the game setting up the Engine and loading levels
func InitGame() *Engine {
	ge := Engine{
		GameState: &sokoban.GameState{
			IsRunning:         false,
			ShowSplash:        false,
			ShowLevel:         false,
			ShowLevelComplete: false,
		},
		Game: &sokoban.Game{
			Levels: map[int]*sokoban.Level{},
		},
	}

	ge.loadLevels()
	// Let's start from the first level.
	ge.loadLevel(1)

	if ge.GUI == nil {
		ge.GUI = &ui.GUI{}
	}
	ge.GUI.Init()

	return &ge
}

// MainLoop is, sic, the main loop of the game
func MainLoop(ge *Engine) {
	ge.GameState.IsRunning = true
	for ge.GameState.IsRunning {
		ge.ManageInput()
		ge.ManageGame()
	}
}
