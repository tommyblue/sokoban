package sokoban

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	log "github.com/sirupsen/logrus"
)

func (ge *engine) drawLevel(screen *ebiten.Image) {
	for i, row := range ge.currentLevel.tiles {
		for j, tileID := range row {
			im := ge.getImage(ge.currentLevel, i, j, tileID)
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

func (ge *engine) getImage(level *level, i, j int, tileID tile) *ebiten.Image {
	if tileID == player && (level.currentPlayerPosition.i != i || level.currentPlayerPosition.j != j) {
		// player has moved, return floor tile
		return ge.ui[floor]
	}

	if level.currentPlayerPosition.i == i && level.currentPlayerPosition.j == j {
		// The player has moved here
		return ge.ui[player]
	}
	r, ok := ge.ui[tileID]
	if !ok {
		return nil
	}
	return r
}

func (ge *engine) loadImages() {
	images := map[tile]string{
		boxOnTarget: "box-ok.png",
		box:         "box.png",
		floor:       "floor.png",
		player:      "player.png",
		target:      "target.png",
		wall:        "wall.png",
	}
	for k, i := range images {
		imgPath := fmt.Sprintf("./assets/%s", i)
		img, _, err := ebitenutil.NewImageFromFile(imgPath, ebiten.FilterDefault)
		if err != nil {
			log.Fatal(err)
		}
		ge.ui[k] = img
	}
}
