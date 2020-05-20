package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/tommyblue/sokoban/game"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	if err := game.Run(); err != nil {
		log.Fatal(err)
	}
}
