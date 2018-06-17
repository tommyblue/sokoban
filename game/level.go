package game

import (
	"bufio"
	"errors"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/tommyblue/sokoban"
	"github.com/tommyblue/sokoban/utils"
)

func (ge *Engine) loadLevel(levelID int) {
	ge.Game.CurrentLevel = &sokoban.Level{}
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
	line string, currentLevelID int, tmpLevel *sokoban.Level,
) (int, *sokoban.Level, error) {
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
		tmpLevel = &sokoban.Level{ID: tmpLevelID}
	} else {
		strTiles := strings.Split(line, "")
		var tiles []sokoban.Tile
		for _, t := range strTiles {
			tiles = append(tiles, sokoban.Tile(t))
		}
		tmpLevel.Tiles = append(tmpLevel.Tiles, tiles)
	}
	return tmpLevelID, tmpLevel, nil
}

func getLevelsFile() (*os.File, func()) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("error")
	}
	filepath := path.Join(path.Dir(filename), "../levels.txt")

	filepath = utils.GetRelativePath("../levels.txt")
	file, err := os.Open(filepath)

	closeFn := func() {
		defer file.Close()
	}

	utils.Check(err)
	return file, closeFn
}
