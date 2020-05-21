package sokoban

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

// Engine represents the game
type engine struct {
	currentLevel *level
	levels       map[int]*level
	ui           map[tile]*ebiten.Image
	movedAt      time.Time
}

func Run() error {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Sokokban")

	ge := &engine{
		levels: map[int]*level{},
		ui:     make(map[tile]*ebiten.Image),
	}
	ge.loadLevels()
	ge.loadLevel(1)

	ge.loadImages()

	return ebiten.RunGame(ge)
}

func (ge *engine) Update(screen *ebiten.Image) error {
	if time.Since(ge.movedAt) > 200*time.Millisecond {
		if ge.checkMovements() {
			ge.movedAt = time.Now()
		}
	}
	return nil
}

func (ge *engine) Draw(screen *ebiten.Image) {
	ge.drawLevel(screen)
}

func (ge *engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
