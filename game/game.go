package game

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 20
	height = 10
)

type Game struct {
	Snake   Snake
	Food    Point
	Score   int
	Running bool
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	snake := Snake{
		Body:  []Point{{width / 2, height / 2}},
		Dir:   Point{0, 1},
		Alive: true,
	}
	game := &Game{
		Snake:   snake,
		Food:    generateFood(),
		Running: true,
	}
	return game
}

func generateFood() Point {
	return Point{rand.Intn(width), rand.Intn(height)}
}

func (g *Game) Update() {
	if !g.Snake.Alive {
		g.Running = false
		return
	}

	g.Snake.Move()

	if g.Snake.Body[0] == g.Food {
		g.Snake.Body = append([]Point{g.Food}, g.Snake.Body...)
		g.Food = generateFood()
		g.Score++
	} else if len(g.Snake.Body) > 1 {
		g.Snake.Body = g.Snake.Body[:len(g.Snake.Body)-1]
	}

	g.Snake.CheckCollision()
}

func (g *Game) HandleInput(dir Point) {
	g.Snake.ChangeDirection(dir)
}

func (g *Game) EndGame() {
	fmt.Println("Game Over! Your score is:", g.Score)
}
