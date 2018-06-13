package game

import "github.com/tommyblue/sokoban"

type direction struct {
	dX int
	dY int
}

var right = direction{dX: 0, dY: 1}
var left = direction{dX: 0, dY: -1}
var up = direction{dX: -1, dY: 0}
var down = direction{dX: 1, dY: 0}

func (ge *GameEngine) MoveLeft() {
	ge.moveIfPossible(left)
}

func (ge *GameEngine) MoveRight() {
	ge.moveIfPossible(right)
}

func (ge *GameEngine) MoveUp() {
	ge.moveIfPossible(up)
}

func (ge *GameEngine) MoveDown() {
	ge.moveIfPossible(down)
}

func (ge *GameEngine) canMoveThere(d direction) bool {
	newI := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI + d.dX
	newJ := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ + d.dY
	tileID := ge.Game.CurrentLevel.Tiles[newI][newJ]

	return tileID != "#"
}

func (ge *GameEngine) move(d direction) {
	currI := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI
	currJ := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ
	ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI = currI + d.dX
	ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ = currJ + d.dY
}

func (ge *GameEngine) moveIfPossible(d direction) *sokoban.PlayerPosition {
	if ge.canMoveThere(d) {
		ge.move(d)
	}
	return ge.Game.CurrentLevel.CurrentPlayerPosition
}
