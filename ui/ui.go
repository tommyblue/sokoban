package ui

import (
	"strings"
	"time"

	"github.com/tommyblue/sokoban"
	"github.com/tommyblue/sokoban/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type GUI struct {
	Timer         time.Time
	PreviousTimer time.Time
	window        *sdl.Window
	renderer      *sdl.Renderer
	countedFrames uint32
	isRunning     bool
	imagesCache   map[string]ImageStruct
}

type ImageStruct struct {
	Name  string
	Image *sdl.Texture
	Rect  sdl.Rect
}

// Init UI components
func (gui *GUI) Init() {
	gui.initSdl()
	gui.initFonts()
	gui.initWindow()
	gui.initRenderer()
	gui.preloadImages()
	gui.isRunning = true
}

// Close ui components
func (gui *GUI) Close() {
	gui.closeRenderer()
	gui.closeWindow()
	gui.closeSdl()
}

func (gui *GUI) Draw(level *sokoban.Level) {
	gui.PreviousTimer = gui.Timer
	gui.Timer = time.Now()
	// updateStatus()
	gui.drawLevel(level)
	gui.finalize()
}

func (gui *GUI) finalize() {

	gui.syncFPS()
	gui.countedFrames++

	gui.renderer.Present()
	gui.renderer.SetDrawColor(167, 125, 83, 255)
	gui.renderer.Clear()
}

func (gui *GUI) drawLevel(level *sokoban.Level) {
	for i, row := range level.Tiles {
		for j, tile := range strings.Split(row, "") {
			image := gui.imagesCache[tile]
			src := image.Rect

			x := imageSide * int32(j)
			y := imageSide * int32(i)
			dst := sdl.Rect{X: x, Y: y, W: imageSide, H: imageSide}
			err := gui.renderer.Copy(image.Image, &src, &dst)
			utils.Check(err)
		}
	}
}
