package sokoban

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	log "github.com/sirupsen/logrus"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct {
	CurrentLevel *Level
	Levels       map[int]*Level
}

type GameState struct {
	ShowSplash        bool
	ShowLevel         bool
	ShowLevelComplete bool
}

// Engine represents the game
type Engine struct {
	GameState *GameState
	Game      *Game
	UI        map[Tile]*ebiten.Image
	movedAt   time.Time
}

type TileDesc struct {
	X int
	Y int
	W int
	H int
}

func (ge *Engine) Update(screen *ebiten.Image) error {
	if time.Since(ge.movedAt) > 200*time.Millisecond {
		if ge.checkMovements() {
			ge.movedAt = time.Now()
		}
	}
	return nil
}

func (ge *Engine) Draw(screen *ebiten.Image) {
	ge.drawLevel(screen)
}

func (ge *Engine) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (ge *Engine) drawLevel(screen *ebiten.Image) {
	for i, row := range ge.Game.CurrentLevel.Tiles {
		for j, tileID := range row {
			im := ge.getImage(ge.Game.CurrentLevel, i, j, tileID)
			if im == nil {
				continue
			}
			w, _ := im.Size()
			x := w * j
			y := w * i

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			if err := screen.DrawImage(im, op); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func (ge *Engine) getImage(level *Level, i, j int, tileID Tile) *ebiten.Image {
	if tileID == Player && (level.CurrentPlayerPosition.PositionI != i || level.CurrentPlayerPosition.PositionJ != j) {
		// player has moved, return floor tile
		return ge.UI[Floor]
	}

	if level.CurrentPlayerPosition.PositionI == i && level.CurrentPlayerPosition.PositionJ == j {
		// The player has moved here
		return ge.UI[Player]
	}
	r, ok := ge.UI[tileID]
	if !ok {
		return nil
	}
	return r
}

func Run() error {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Sokokban")

	ge := &Engine{
		GameState: &GameState{
			ShowSplash:        false,
			ShowLevel:         false,
			ShowLevelComplete: false,
		},
		Game: &Game{
			Levels: map[int]*Level{},
		},
		UI: make(map[Tile]*ebiten.Image),
	}
	ge.loadLevels()
	ge.loadLevel(1)

	ge.loadImages()

	return ebiten.RunGame(ge)
}

func (ge *Engine) loadImages() {
	images := map[Tile]string{
		BoxOnTarget: "box-ok.png",
		Box:         "box.png",
		Floor:       "floor.png",
		Player:      "player.png",
		Target:      "target.png",
		Wall:        "wall.png",
	}
	for k, i := range images {
		imgPath := fmt.Sprintf("./assets/%s", i)
		img, _, err := ebitenutil.NewImageFromFile(imgPath, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		ge.UI[k] = img
	}
}
