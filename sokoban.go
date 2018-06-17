package sokoban

import (
	"fmt"

	"github.com/tommyblue/sokoban/utils"
)

type Tile string

const (
	Wall        Tile = "#"
	Target      Tile = "."
	Floor       Tile = "_"
	Box         Tile = "$"
	BoxOnTarget Tile = "+"
	Empty       Tile = "~"
	Player      Tile = "@"
)

// Level describes a level of the game
type Level struct {
	ID                    int
	Width                 int
	Height                int
	Tiles                 [][]Tile
	CurrentPlayerPosition *PlayerPosition
	TilesToFix            int
}

// GameState is the state of the game
type GameState struct {
	IsRunning         bool
	ShowSplash        bool
	ShowLevel         bool
	ShowLevelComplete bool
}

type PlayerPosition struct {
	PositionI int
	PositionJ int
}

type Game struct {
	CurrentLevel *Level
	Levels       map[int]*Level
}

func (l *Level) CloneFrom(orig *Level) {
	l.ID = orig.ID
	l.Width = orig.Width
	l.Height = orig.Height
	l.TilesToFix = orig.TilesToFix
	l.CurrentPlayerPosition = &PlayerPosition{
		PositionI: orig.CurrentPlayerPosition.PositionI,
		PositionJ: orig.CurrentPlayerPosition.PositionJ,
	}
	for _, row := range orig.Tiles {
		var tiles []Tile
		for _, tile := range row {
			tiles = append(tiles, tile)
		}
		l.Tiles = append(l.Tiles, tiles)
	}
}

func (l *Level) Finalize() {
	h, w := 0, 0
	for i, row := range l.Tiles {
		h = i + 1
		if w < len(row) {
			w = len(row)
		}
		for j, tile := range row {
			if tile == Player {
				l.CurrentPlayerPosition = &PlayerPosition{
					PositionI: i,
					PositionJ: j,
				}
			}
			if tile == Target {
				l.TilesToFix++
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
