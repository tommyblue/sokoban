package game

func (ge *GameEngine) MoveLeft() {
	i := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI
	j := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ
	if ge.canMoveThere(i, j-1) {
		ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ = j - 1
	}
}

func (ge *GameEngine) MoveRight() {
	i := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI
	j := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ
	if ge.canMoveThere(i, j+1) {
		ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ = j + 1
	}
}

func (ge *GameEngine) MoveUp() {
	i := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI
	j := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ
	if ge.canMoveThere(i-1, j) {
		ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI = i - 1
	}
}

func (ge *GameEngine) MoveDown() {
	i := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI
	j := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ
	if ge.canMoveThere(i+1, j) {
		ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI = i + 1
	}
}

func (ge *GameEngine) canMoveThere(i, j int) bool {
	tileID := ge.Game.CurrentLevel.Tiles[i][j]

	return tileID != "#"
}
