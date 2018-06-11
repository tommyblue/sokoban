package ui

import (
	"github.com/tommyblue/sokoban/utils"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func (gui *GUI) initSdl() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	utils.Check(err)
}

func (gui *GUI) closeSdl() {
	sdl.Quit()
}

func (gui *GUI) initWindow() {
	var err error
	gui.window, err = sdl.CreateWindow(
		"Sokoban",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800,
		600,
		sdl.WINDOW_SHOWN,
	)
	utils.Check(err)
}

func (gui *GUI) closeWindow() {
	gui.window.Destroy()
}

func (gui *GUI) initRenderer() {
	var err error
	gui.renderer, err = sdl.CreateRenderer(gui.window, -1, sdl.RENDERER_ACCELERATED)
	utils.Check(err)
	gui.renderer.SetDrawColor(0, 255, 0, 255)
	gui.renderer.Clear()
}

func (gui *GUI) closeRenderer() {
	gui.renderer.Destroy()
}

func (gui *GUI) initFonts() {
	// Initialize TTF
	ttf.Init()
	var err error
	filepath := utils.GetRelativePath("../assets/fonts/mono.ttf")
	font, err = ttf.OpenFont(filepath, 14)
	utils.Check(err)
}
