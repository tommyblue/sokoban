package ui

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type GUI struct {
	Timer         time.Time
	PreviousTimer time.Time
	window        *sdl.Window
	renderer      *sdl.Renderer
	countedFrames uint32
	isRunning     bool
}

// Init UI components
func (gui *GUI) Init() {
	gui.initSdl()
	gui.initFonts()
	gui.initWindow()
	gui.initRenderer()
	gui.isRunning = true
}

// Close ui components
func (gui *GUI) Close() {
	gui.closeRenderer()
	gui.closeWindow()
	gui.closeSdl()
}

func (gui *GUI) Loop() {
	gui.Timer = time.Now()
	for gui.isRunning {
		gui.PreviousTimer = gui.Timer
		gui.Timer = time.Now()
		gui.manageInput()
		// updateStatus()
		gui.render()
	}
}

func (gui *GUI) manageInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			gui.isRunning = false
			break
		}
	}
}

func (gui *GUI) render() {
	gui.syncFPS()
	gui.countedFrames++

	gui.renderer.Present()
	gui.renderer.SetDrawColor(167, 125, 83, 255)
	gui.renderer.Clear()
}
