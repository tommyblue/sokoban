package sokoban

import (
	"fmt"

	"github.com/tommyblue/sokoban/utils"
)

// Level describes a level of the game
type Level struct {
	ID                    int
	Width                 int
	Height                int
	Tiles                 [][]string
	CurrentPlayerPosition *PlayerPosition
}

type PlayerPosition struct {
	PositionI int
	PositionJ int
}

type Game struct {
	CurrentLevel *Level
	Levels       []*Level
}

func (l *Level) Finalize() {
	h, w := 0, 0
	for i, row := range l.Tiles {
		h = i + 1
		if w < len(row) {
			w = len(row)
		}
		for j, tile := range row {
			if tile == "@" {
				l.CurrentPlayerPosition = &PlayerPosition{
					PositionI: i,
					PositionJ: j,
				}
			}
		}
	}
	l.Height = h
	l.Width = w

	if utils.IsDebugEnv() {
		l.printInfo()
	}
}

func (l *Level) printInfo() {
	fmt.Printf("ID: %d\n", l.ID)
	fmt.Printf("Size: %dx%d\n", l.Width, l.Height)
}
