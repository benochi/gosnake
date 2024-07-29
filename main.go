package main

import (
	"fmt"
	"snake-game/game"
	"snake-game/render"

	"github.com/faiface/pixel/pixelgl"
)

func main() {
	fmt.Println("Starting game...")
	pixelgl.Run(run)
}

func run() {
	fmt.Println("Running game...")
	game := game.NewGame()
	render.Run(game)
}
