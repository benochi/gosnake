package render

import (
	"fmt"
	"gosnake/constants"
	"gosnake/game"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func Run(g *game.Game) {
	fmt.Println("Initializing window...")
	cfg := pixelgl.WindowConfig{
		Title:  "Snake Game",
		Bounds: pixel.R(0, 0, constants.WinWidth, constants.WinHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("Entering game loop...")
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for !win.Closed() {
		select {
		case <-ticker.C:
			if g.Running {
				fmt.Println("Updating game...")
				g.Update()
			}
		default:
			handleInput(win, g)
			win.Clear(colornames.Black)
			drawGame(win, g)
			win.Update()
		}

		if win.Pressed(pixelgl.KeyQ) {
			win.SetClosed(true)
		}
	}
}

func handleInput(win *pixelgl.Window, g *game.Game) {
	if win.JustPressed(pixelgl.KeyW) {
		g.HandleInput(game.Point{0, 1}) // Fix: W moves up
	}
	if win.JustPressed(pixelgl.KeyA) {
		g.HandleInput(game.Point{-1, 0})
	}
	if win.JustPressed(pixelgl.KeyS) {
		g.HandleInput(game.Point{0, -1}) // Fix: S moves down
	}
	if win.JustPressed(pixelgl.KeyD) {
		g.HandleInput(game.Point{1, 0})
	}
}

func drawGame(win *pixelgl.Window, g *game.Game) {
	imd := imdraw.New(nil)

	imd.Color = colornames.Green
	for _, p := range g.Snake.Body {
		imd.Push(pixel.V(float64(p.X*constants.CellSize), float64(p.Y*constants.CellSize)))
		imd.Push(pixel.V(float64(p.X*constants.CellSize+constants.CellSize), float64(p.Y*constants.CellSize+constants.CellSize)))
		imd.Rectangle(0)
	}

	imd.Color = colornames.Red
	imd.Push(pixel.V(float64(g.Food.X*constants.CellSize), float64(g.Food.Y*constants.CellSize)))
	imd.Push(pixel.V(float64(g.Food.X*constants.CellSize+constants.CellSize), float64(g.Food.Y*constants.CellSize+constants.CellSize)))
	imd.Rectangle(0)

	imd.Draw(win)
}
