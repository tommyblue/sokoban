package ui

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tommyblue/sokoban/utils"
	"github.com/veandco/go-sdl2/sdl"
)

func (gui *GUI) syncFPS() {
	tick := time.Now()
	elapsedMS := float64(tick.Sub(gui.Timer)) / float64(time.Millisecond)
	if utils.IsDebugEnv() {
		var zeroTime time.Time
		if lastTimestamp == zeroTime {
			lastTimestamp = time.Now()
		} else {
			if time.Since(lastTimestamp) < time.Duration(time.Second) {
				countSinceLast++
			} else {
				lastTimestamp = time.Now()
				if os.Getenv("FPSLOG") == "1" {
					log.Printf("[%v] update FPS: %v -> %v\n", lastTimestamp, lastFPS, countSinceLast+1)
				}
				lastFPS = countSinceLast + 1
				countSinceLast = 0
			}
			gui.drawFPS(fmt.Sprintf("%v", lastFPS))
		}
	}
	if sleep := TICKSPERFRAME - elapsedMS; sleep > 0 {
		d := time.Duration(sleep)
		sdl.Delay(uint32(d))
	}
}

func (gui *GUI) drawFPS(text string) error {
	fontColor := sdl.Color{R: 0, G: 0, B: 0, A: 0}
	textSurface, err := font.RenderUTF8_Blended(fmt.Sprintf("%s FPS", text), fontColor)
	utils.Check(err)
	defer textSurface.Free()

	textTexture, err := gui.renderer.CreateTextureFromSurface(textSurface)
	if err != nil {
		return err
	}
	defer textTexture.Destroy()

	src := sdl.Rect{X: 0, Y: 0, W: textSurface.W, H: textSurface.H}
	dest := sdl.Rect{X: 0, Y: 0, W: textSurface.W, H: textSurface.H}
	gui.renderer.Copy(textTexture, &src, &dest)

	return nil
}
