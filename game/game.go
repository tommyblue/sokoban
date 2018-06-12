package game

import (
	"bufio"
	"errors"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/tommyblue/sokoban"
	"github.com/tommyblue/sokoban/ui"
	"github.com/tommyblue/sokoban/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type GameEngine struct {
	IsRunning bool
	Game      *sokoban.Game
	// TODO: Remove ui dependency (move to sokoban?)
	GUI *ui.GUI
}

// InitGame initializes the game
func InitGame() *GameEngine {
	ge := GameEngine{
		Game: &sokoban.Game{
			Levels: []*sokoban.Level{},
		},
	}

	ge.loadLevels()
	// Let's start from the first level
	ge.Game.CurrentLevel = ge.Game.Levels[1]
	return &ge
}

func (ge *GameEngine) ManageInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			ge.IsRunning = false
			break
		}
	}
}

func (ge *GameEngine) loadLevels() {
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

func (ge *GameEngine) parseLevelString(
	line string, currentLevelID int, tmpLevel *sokoban.Level,
) (int, *sokoban.Level, error) {
	tmpLevelID := currentLevelID
	if line == ";END" {
		tmpLevel.CalculateSize()
		ge.Game.Levels = append(ge.Game.Levels, tmpLevel)
		return tmpLevelID, tmpLevel, errors.New("Reached end of file")
	}

	if strings.HasPrefix(line, ";LEVEL") {
		// Do not add an empty level if this is the first one
		if tmpLevel != nil {
			tmpLevel.CalculateSize()
			ge.Game.Levels = append(ge.Game.Levels, tmpLevel)
		}
		tmpLevelID++
		tmpLevel = &sokoban.Level{ID: tmpLevelID}
	} else {
		tmpLevel.Tiles = append(tmpLevel.Tiles, line)
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
