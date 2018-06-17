package game

import (
	"fmt"

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

// MoveLeft moves... left
func (ge *Engine) MoveLeft() {
	ge.moveIfPossible(left)
}

// MoveRight moves... right
func (ge *Engine) MoveRight() {
	ge.moveIfPossible(right)
}

// MoveUp moves... up
func (ge *Engine) MoveUp() {
	ge.moveIfPossible(up)
}

// MoveDown moves... down
func (ge *Engine) MoveDown() {
	ge.moveIfPossible(down)
}

func (ge *Engine) canMoveThere(d direction) bool {
	newI := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI + d.dX
	newJ := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ + d.dY
	tileID := ge.Game.CurrentLevel.Tiles[newI][newJ]

	// Can't move through walls
	if tileID == sokoban.Wall {
		return false
	}
	// Can move to box if next tile isn't box or wall
	if tileID == sokoban.Box || tileID == sokoban.BoxOnTarget {
		nextTile := ge.Game.CurrentLevel.Tiles[newI+d.dX][newJ+d.dY]
		return nextTile != sokoban.Box && nextTile != sokoban.BoxOnTarget && nextTile != sokoban.Wall
	}
	return true
}

func (ge *Engine) move(d direction) {
	currI := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI
	currJ := ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ
	nextI := currI + d.dX
	nextJ := currJ + d.dY
	ge.Game.CurrentLevel.CurrentPlayerPosition.PositionI = nextI
	ge.Game.CurrentLevel.CurrentPlayerPosition.PositionJ = nextJ

	// Check if a box must be moved
	if ge.Game.CurrentLevel.Tiles[nextI][nextJ] == sokoban.Box ||
		ge.Game.CurrentLevel.Tiles[nextI][nextJ] == sokoban.BoxOnTarget {
		replaceTile := sokoban.Floor
		if ge.Game.Levels[ge.Game.CurrentLevel.ID].Tiles[nextI][nextJ] == sokoban.Target {
			replaceTile = sokoban.Target
			ge.Game.CurrentLevel.TilesToFix++
		}
		ge.Game.CurrentLevel.Tiles[nextI][nextJ] = replaceTile
		boxTile := sokoban.Box
		if ge.Game.Levels[ge.Game.CurrentLevel.ID].Tiles[nextI+d.dX][nextJ+d.dY] == sokoban.Target {
			boxTile = sokoban.BoxOnTarget
			ge.Game.CurrentLevel.TilesToFix--
		}
		ge.Game.CurrentLevel.Tiles[nextI+d.dX][nextJ+d.dY] = boxTile
	}
}

func (ge *Engine) checkVictory() {
	if ge.Game.CurrentLevel.TilesToFix == 0 {
		fmt.Println("Victory!")
		// Move to next level
		ge.loadLevel(ge.Game.CurrentLevel.ID + 1)
	}
}

func (ge *Engine) moveIfPossible(d direction) *sokoban.PlayerPosition {
	if ge.canMoveThere(d) {
		ge.move(d)
		ge.checkVictory()
	}
	return ge.Game.CurrentLevel.CurrentPlayerPosition
}
