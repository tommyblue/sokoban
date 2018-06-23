package ui

import "github.com/veandco/go-sdl2/sdl"

type IGameEngine interface {
	MoveLeft()
	MoveRight()
	MoveUp()
	MoveDown()
	SetRunningState(bool)
}

// ManageInput checks input from the user (mainly keyboard)
func (gui *GUI) ManageInput(ge IGameEngine) {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.KeyUpEvent:
			switch t.Keysym.Scancode {
			case sdl.SCANCODE_ESCAPE:
				println("Quit")
				ge.SetRunningState(false)
				break
			case sdl.SCANCODE_LEFT:
				ge.MoveLeft()
			case sdl.SCANCODE_RIGHT:
				ge.MoveRight()
			case sdl.SCANCODE_UP:
				ge.MoveUp()
			case sdl.SCANCODE_DOWN:
				ge.MoveDown()
			}
		case *sdl.QuitEvent:
			println("Quit")
			ge.SetRunningState(false)
			break
		}
	}
}
