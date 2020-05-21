package sokoban

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

func (ge *engine) checkMovements() bool {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		ge.moveIfPossible(up)
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		ge.moveIfPossible(down)
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		ge.moveIfPossible(right)
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		ge.moveIfPossible(left)
	} else {
		return false
	}
	return true
}

func (ge *engine) moveIfPossible(d direction) *playerPosition {
	if ge.canMoveThere(d) {
		ge.move(d)
		ge.checkVictory()
	}
	return ge.currentLevel.currentPlayerPosition
}

func (ge *engine) canMoveThere(d direction) bool {
	p := ge.currentLevel.currentPlayerPosition
	nextTile := ge.currentLevel.tiles[p.i+d.dX][p.j+d.dY]

	// Can't move through walls
	if nextTile == wall {
		return false
	}
	// Can move a box if the tile after the box isn't another box or a wall
	if nextTile == box || nextTile == boxOnTarget {
		nextTile := ge.currentLevel.tiles[p.i+d.dX+d.dX][p.j+d.dY+d.dY]
		return nextTile != box && nextTile != boxOnTarget && nextTile != wall
	}
	// In any other case, the player can move
	return true
}

func (ge *engine) move(d direction) {
	nextI := ge.currentLevel.currentPlayerPosition.i + d.dX
	nextJ := ge.currentLevel.currentPlayerPosition.j + d.dY
	ge.currentLevel.currentPlayerPosition.i = nextI
	ge.currentLevel.currentPlayerPosition.j = nextJ

	// Check if a box must be moved
	if ge.currentLevel.tiles[nextI][nextJ] == box || ge.currentLevel.tiles[nextI][nextJ] == boxOnTarget {
		replaceTile := floor
		if ge.levels[ge.currentLevel.id].tiles[nextI][nextJ] == target {
			replaceTile = target
			ge.currentLevel.tilesToFix++
		}
		ge.currentLevel.tiles[nextI][nextJ] = replaceTile
		boxTile := box
		if ge.levels[ge.currentLevel.id].tiles[nextI+d.dX][nextJ+d.dY] == target {
			boxTile = boxOnTarget
			ge.currentLevel.tilesToFix--
		}
		ge.currentLevel.tiles[nextI+d.dX][nextJ+d.dY] = boxTile
	}
}

func (ge *engine) checkVictory() {
	if ge.currentLevel.tilesToFix == 0 {
		fmt.Println("Victory!")
		// Move to next level
		ge.loadLevel(ge.currentLevel.id + 1)
	}
}
