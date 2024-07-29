package render

import (
	"fmt"
	"snake-game/game"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	winWidth  = 800
	winHeight = 600
	cellSize  = 20
)

func Run(g *game.Game) {
	fmt.Println("Initializing window...")
	cfg := pixelgl.WindowConfig{
		Title:  "Snake Game",
		Bounds: pixel.R(0, 0, winWidth, winHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println("Entering game loop...")
	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		if g.Running {
			g.Update()
		}

		win.Clear(colornames.Black)
		drawGame(win, g)
		win.Update()

		if win.Pressed(pixelgl.KeyQ) {
			win.SetClosed(true)
		}

		handleInput(win, g)
		time.Sleep(time.Duration(dt*1000) * time.Millisecond)
	}
}

func handleInput(win *pixelgl.Window, g *game.Game) {
	if win.Pressed(pixelgl.KeyW) {
		g.HandleInput(game.Point{0, -1})
	}
	if win.Pressed(pixelgl.KeyA) {
		g.HandleInput(game.Point{-1, 0})
	}
	if win.Pressed(pixelgl.KeyS) {
		g.HandleInput(game.Point{0, 1})
	}
	if win.Pressed(pixelgl.KeyD) {
		g.HandleInput(game.Point{1, 0})
	}
}

func drawGame(win *pixelgl.Window, g *game.Game) {
	imd := imdraw.New(nil)

	imd.Color = colornames.Green
	for _, p := range g.Snake.Body {
		imd.Push(pixel.V(float64(p.X*cellSize), float64(p.Y*cellSize)))
		imd.Push(pixel.V(float64(p.X*cellSize+cellSize), float64(p.Y*cellSize+cellSize)))
		imd.Rectangle(0)
	}

	imd.Color = colornames.Red
	imd.Push(pixel.V(float64(g.Food.X*cellSize), float64(g.Food.Y*cellSize)))
	imd.Push(pixel.V(float64(g.Food.X*cellSize+cellSize), float64(g.Food.Y*cellSize+cellSize)))
	imd.Rectangle(0)

	imd.Draw(win)
}
