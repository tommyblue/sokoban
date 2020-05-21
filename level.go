package sokoban

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/tommyblue/sokoban/utils"
)

type Level struct {
	ID                    int
	Width                 int
	Height                int
	Tiles                 [][]Tile
	CurrentPlayerPosition *PlayerPosition
	TilesToFix            int
}

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

type PlayerPosition struct {
	PositionI int
	PositionJ int
}

func (ge *Engine) loadLevel(levelID int) {
	ge.Game.CurrentLevel = &Level{}
	ge.Game.CurrentLevel.CloneFrom(ge.Game.Levels[levelID])
}

func (ge *Engine) loadLevels() {
	file, closeFn := getLevelsFile()
	defer closeFn()
	levelID := 0
	for {
		r := bufio.NewReader(file)
		line, err := utils.ReadLine(r)
		levelID, tmpLevel, e := ge.parseLevelString(line, levelID, nil)
		if e != nil {
			break
		}
		for err == nil {
			line, err = utils.ReadLine(r)
			levelID, tmpLevel, e = ge.parseLevelString(line, levelID, tmpLevel)
			if e != nil {
				break
			}
		}
		break
	}
}

func (ge *Engine) parseLevelString(
	line string, currentLevelID int, tmpLevel *Level,
) (int, *Level, error) {
	tmpLevelID := currentLevelID
	if line == ";END" {
		tmpLevel.Finalize()
		ge.Game.Levels[tmpLevel.ID] = tmpLevel
		return tmpLevelID, tmpLevel, errors.New("Reached end of file")
	}

	if strings.HasPrefix(line, ";LEVEL") {
		// Do not add an empty level if this is the first one
		if tmpLevel != nil {
			tmpLevel.Finalize()
			ge.Game.Levels[tmpLevel.ID] = tmpLevel
		}
		tmpLevelID++
		tmpLevel = &Level{ID: tmpLevelID}
	} else {
		strTiles := strings.Split(line, "")
		var tiles []Tile
		for _, t := range strTiles {
			tiles = append(tiles, Tile(t))
		}
		tmpLevel.Tiles = append(tmpLevel.Tiles, tiles)
	}
	return tmpLevelID, tmpLevel, nil
}

func getLevelsFile() (*os.File, func()) {
	filepath := utils.GetRelativePath("./levels.txt")
	file, err := os.Open(filepath)

	closeFn := func() {
		defer file.Close()
	}

	utils.Check(err)
	return file, closeFn
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
		tiles = append(tiles, row...)
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
