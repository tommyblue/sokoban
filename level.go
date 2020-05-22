package sokoban

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/tommyblue/sokoban/utils"
)

type level struct {
	id                    int
	width                 int
	height                int
	tiles                 [][]tile
	currentPlayerPosition *playerPosition
	tilesToFix            int
}

type tile string

const (
	wall        tile = "#"
	target      tile = "."
	floor       tile = "_"
	box         tile = "$"
	boxOnTarget tile = "+"
	player      tile = "@"
)

type playerPosition struct {
	i int
	j int
}

func (ge *engine) loadLevel(levelID int) {
	ge.currentLevel = &level{}
	ge.currentLevel.cloneFrom(ge.levels[levelID])
}

func (ge *engine) loadLevels() {
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

func (ge *engine) parseLevelString(
	line string, currentLevelID int, tmpLevel *level,
) (int, *level, error) {
	tmpLevelID := currentLevelID
	if line == ";END" {
		tmpLevel.finalize()
		ge.levels[tmpLevel.id] = tmpLevel
		return tmpLevelID, tmpLevel, errors.New("Reached end of file")
	}

	if strings.HasPrefix(line, ";LEVEL") {
		// Do not add an empty level if this is the first one
		if tmpLevel != nil {
			tmpLevel.finalize()
			ge.levels[tmpLevel.id] = tmpLevel
		}
		tmpLevelID++
		tmpLevel = &level{id: tmpLevelID}
	} else {
		strTiles := strings.Split(line, "")
		var tiles []tile
		for _, t := range strTiles {
			tiles = append(tiles, tile(t))
		}
		tmpLevel.tiles = append(tmpLevel.tiles, tiles)
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

func (l *level) cloneFrom(orig *level) {
	l.id = orig.id
	l.width = orig.width
	l.height = orig.height
	l.tilesToFix = orig.tilesToFix
	l.currentPlayerPosition = &playerPosition{
		i: orig.currentPlayerPosition.i,
		j: orig.currentPlayerPosition.j,
	}
	for _, row := range orig.tiles {
		var tiles []tile
		tiles = append(tiles, row...)
		l.tiles = append(l.tiles, tiles)
	}
}

func (l *level) finalize() {
	h, w := 0, 0
	for i, row := range l.tiles {
		h = i + 1
		if w < len(row) {
			w = len(row)
		}
		for j, tile := range row {
			if tile == player {
				l.currentPlayerPosition = &playerPosition{
					i: i,
					j: j,
				}
			}
			if tile == target {
				l.tilesToFix++
			}
		}
	}
	l.height = h
	l.width = w
}
