package game

import (
	"github.com/tommyblue/sokoban"
)

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

	// Can't move through walls
	if tileID == sokoban.Wall {
		return false
	}
	// Can move to box if next tile isn't box or wall
	if tileID == sokoban.Box {
		nextTile := ge.Game.CurrentLevel.Tiles[newI+d.dX][newJ+d.dY]
		return nextTile != sokoban.Box && nextTile != sokoban.Wall
	}
	return true
}

func (ge *GameEngine) move(d direction) {
	currI := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI
	currJ := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ
	nextI := currI + d.dX
	nextJ := currJ + d.dY
	ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI = nextI
	ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ = nextJ

	// Check if a box must be moved
	if ge.Game.CurrentLevel.Tiles[nextI][nextJ] == sokoban.Box {
		ge.Game.CurrentLevel.Tiles[nextI][nextJ] = sokoban.Floor
		ge.Game.CurrentLevel.Tiles[nextI+d.dX][nextJ+d.dY] = sokoban.Box
	}
}

func (ge *GameEngine) moveIfPossible(d direction) *sokoban.PlayerPosition {
	if ge.canMoveThere(d) {
		ge.move(d)
	}
	return ge.Game.CurrentLevel.CurrentPlayerPosition
}
