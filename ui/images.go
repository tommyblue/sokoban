package ui

import (
	"github.com/tommyblue/sokoban"
	"github.com/tommyblue/sokoban/utils"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func (gui *GUI) preloadImages() {
	if gui.imagesCache == nil {
		gui.imagesCache = make(map[sokoban.Tile]ImageStruct)
	}
	tiles := map[sokoban.Tile]string{
		sokoban.Wall:        "wall",
		sokoban.Target:      "target",
		sokoban.Floor:       "floor",
		sokoban.Box:         "box",
		sokoban.BoxOnTarget: "box-ok",
		sokoban.Empty:       "empty",
		sokoban.Player:      "man-bottom",
	}
	for tileID, tileName := range tiles {
		surface, err := sdl.CreateRGBSurface(
			0,
			imageSide,
			imageSide,
			32,
			0xff000000,
			0x00ff0000,
			0x0000ff00,
			0x000000ff,
		)
		utils.Check(err)

		srcRect := sdl.Rect{X: 0, Y: 0, W: imageSide, H: imageSide}

		if tileID != "~" {
			if tileID == "@" {
				floor, err := img.Load(utils.GetRelativePath("../assets/images/floor.png"))
				utils.Check(err)

				image, err := img.Load(utils.GetRelativePath("../assets/images/" + tileName + ".png"))
				utils.Check(err)

				err = floor.Blit(&srcRect, surface, &srcRect)
				utils.Check(err)

				err = image.Blit(&srcRect, surface, &srcRect)
				utils.Check(err)
			} else {
				image, err := img.Load(utils.GetRelativePath("../assets/images/" + tileName + ".png"))
				utils.Check(err)
				err = image.Blit(&srcRect, surface, &srcRect)
				utils.Check(err)
			}
		}

		imgTexture, err := gui.renderer.CreateTextureFromSurface(surface)
		utils.Check(err)

		imgStruct := ImageStruct{Name: tileName, Image: imgTexture, Rect: srcRect}
		gui.imagesCache[tileID] = imgStruct
	}
}
