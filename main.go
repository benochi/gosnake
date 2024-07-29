package main

import (
	"fmt"
	"gosnake/game"
	"gosnake/render"

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
