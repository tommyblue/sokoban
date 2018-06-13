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
			Levels: map[int]*sokoban.Level{},
		},
	}

	ge.loadLevels()
	// Let's start from the first level.
	// Make a copy so that the original level is always available
	ge.Game.CurrentLevel = &sokoban.Level{}
	ge.Game.CurrentLevel.CloneFrom(ge.Game.Levels[1])

	return &ge
}

func (ge *GameEngine) ManageInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.KeyUpEvent:
			switch t.Keysym.Scancode {
			case sdl.SCANCODE_ESCAPE:
				println("Quit")
				ge.IsRunning = false
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
