package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/tommyblue/sokoban"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	if err := sokoban.Run(); err != nil {
		log.Fatal(err)
	}
}
