package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

type direction struct {
	dX int
	dY int
}

var (
	right = direction{dX: 0, dY: 1}
	left  = direction{dX: 0, dY: -1}
	up    = direction{dX: -1, dY: 0}
	down  = direction{dX: 1, dY: 0}
)

func (ge *Engine) checkMovements() bool {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ge.MoveUp()
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ge.MoveDown()
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ge.MoveRight()
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ge.MoveLeft()
	} else {
		return false
	}
	return true
}

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
	if tileID == Wall {
		return false
	}
	// Can move to box if next tile isn't box or wall
	if tileID == Box || tileID == BoxOnTarget {
		nextTile := ge.Game.CurrentLevel.Tiles[newI+d.dX][newJ+d.dY]
		return nextTile != Box && nextTile != BoxOnTarget && nextTile != Wall
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
	if ge.Game.CurrentLevel.Tiles[nextI][nextJ] == Box ||
		ge.Game.CurrentLevel.Tiles[nextI][nextJ] == BoxOnTarget {
		replaceTile := Floor
		if ge.Game.Levels[ge.Game.CurrentLevel.ID].Tiles[nextI][nextJ] == Target {
			replaceTile = Target
			ge.Game.CurrentLevel.TilesToFix++
		}
		ge.Game.CurrentLevel.Tiles[nextI][nextJ] = replaceTile
		boxTile := Box
		if ge.Game.Levels[ge.Game.CurrentLevel.ID].Tiles[nextI+d.dX][nextJ+d.dY] == Target {
			boxTile = BoxOnTarget
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

func (ge *Engine) moveIfPossible(d direction) *PlayerPosition {
	if ge.canMoveThere(d) {
		ge.move(d)
		ge.checkVictory()
	}
	return ge.Game.CurrentLevel.CurrentPlayerPosition
}
