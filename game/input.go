package game

import "github.com/veandco/go-sdl2/sdl"

// ManageInput checks input from the user (mainly keyboard)
func (ge *Engine) ManageInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.KeyUpEvent:
			switch t.Keysym.Scancode {
			case sdl.SCANCODE_ESCAPE:
				println("Quit")
				ge.GameState.IsRunning = false
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
			ge.GameState.IsRunning = false
			break
		}
	}
}
