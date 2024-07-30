package game

import (
	"fmt"
	"gosnake/constants"
	"math/rand"
	"time"
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
		Body:  []Point{{constants.Width / 2, constants.Height / 2}},
		Dir:   Point{0, 1},
		Alive: true,
	}
	game := &Game{
		Snake:   snake,
		Food:    generateFood(),
		Running: true,
	}
	fmt.Println("Game initialized:", game)
	return game
}

func generateFood() Point {
	return Point{rand.Intn(constants.Width), rand.Intn(constants.Height)}
}

func (g *Game) Update() {
	if !g.Snake.Alive {
		g.Running = false
		fmt.Println("Snake is not alive. Stopping game.")
		return
	}

	// Save the current length of the snake
	initialLength := len(g.Snake.Body)

	// Move the snake
	g.Snake.Move()

	// Check if food is eaten
	if g.Snake.Body[0] == g.Food {
		g.Food = generateFood()
		g.Score++
		fmt.Println("Food eaten. New food generated.")
	} else {
		// Remove the tail if food is not eaten
		if len(g.Snake.Body) > initialLength {
			g.Snake.Body = g.Snake.Body[:initialLength]
		}
	}
	g.Snake.CheckCollision()
}

func (g *Game) HandleInput(dir Point) {
	g.Snake.ChangeDirection(dir)
}

func (g *Game) EndGame() {
	fmt.Println("Game Over! Your score is:", g.Score)
}
