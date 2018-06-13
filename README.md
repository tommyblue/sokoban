# Sokoban

A Go implementation of the famous puzzle game [Sokoban](http://www.sokoban.jp/) by Hiroyuki
Imabayashi.

## Levels

The levels are stored as text file in the `assets/levels/` folder.
Each file contains the matrix of the level where each tile can have one of these values:

- `@`: player
- `#`: wall
- `$`: box
- `.`: target (where to move the box)
- `+`: box on a target
- `~`: empty space, where the player/boxes can be moved

## Assets

Assets come from [Keeney Sokoban package](http://kenney.nl/assets/sokoban)
distributed with [CC0 1.0 Universal (CC0 1.0)](https://creativecommons.org/publicdomain/zero/1.0/)
license.

## Setup

You need SDL2 installed. You can find the instructions for supported systems in the [Go-SDL2 page](https://github.com/veandco/go-sdl2)

## Compile and run

`DEBUG=1 go run cmd/sokoban/main.go`
