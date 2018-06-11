package sokoban

import (
	"fmt"
)

// Level describes a level of the game
type Level struct {
	ID     int
	Width  int
	Height int
	Tiles  []string
}

func (l *Level) CalculateSize() {
	h, w := 0, 0
	for i, line := range l.Tiles {
		h = i + 1
		if w < len(line) {
			w = len(line)
		}
	}
	l.Height = h
	l.Width = w
	l.PrintInfo()
}

func (l *Level) PrintInfo() {
	fmt.Printf("ID: %d\n", l.ID)
	fmt.Printf("Size: %dx%d\n", l.Width, l.Height)
}

type Game struct {
	CurrentLevel *Level
	Levels       []*Level
}

type GameEngine interface {
	LoadLevels() *[]Level
}
