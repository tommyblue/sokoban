package game

// ManageGame checks the game status and manages it, e.g. decides what to show the user
// (the game, some stats, etc.)
func (ge *Engine) ManageGame() {

	/* Check game status and decide whether to show:
	- intro
	- level
	- level-completed modal
	*/
	ge.GUI.Draw(ge.Game.CurrentLevel)
}
