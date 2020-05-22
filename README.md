# Sokoban

A Go implementation of the famous puzzle game [Sokoban](http://www.sokoban.jp/) by Hiroyuki
Imabayashi using the [Ebiten](https://ebiten.org/) 2D game library.

![Alt text](screenshot.jpg?raw=true "Screenshot")

## Levels

The levels are stored in the [levels.txt](./levels.txt) file.
Each section contains the matrix of the level where each tile can have one of these values:

- `@`: player
- `#`: wall
- `$`: box
- `.`: target (where to move the box)
- `+`: box on a target
- `_`: empty space, where the player/boxes can be moved

## Assets

Assets come from [Keeney Sokoban package](http://kenney.nl/assets/sokoban)
distributed with [CC0 1.0 Universal (CC0 1.0)](https://creativecommons.org/publicdomain/zero/1.0/)
license.

## Setup

Ebiten has some requirements, please take a look at its [install page](https://ebiten.org/documents/install.html)

## Make commands

This is the output of `make help`:

```
build                          Build binary in the local env
govet                          Run go vet on the project
run                            Run the app
test                           Run go tests
```
