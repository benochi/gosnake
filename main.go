package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	width  = 20
	height = 10
)

type Point struct {
	x, y int
}

type Snake struct {
	body  []Point
	dir   Point
	alive bool
}

type Game struct {
	snake   Snake
	food    Point
	score   int
	running bool
}

func main() {
	game := NewGame()
	game.Run()
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	snake := Snake{
		body:  []Point{{width / 2, height / 2}},
		dir:   Point{0, 1},
		alive: true,
	}
	game := &Game{
		snake:   snake,
		food:    generateFood(),
		running: true,
	}
	return game
}

func generateFood() Point {
	return Point{rand.Intn(width), rand.Intn(height)}
}

func (g *Game) Run() {
	go g.readInput()

	for g.running {
		g.Update()
		g.Draw()
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("Game Over! Your score is: ", g.score)
}

func (g *Game) readInput() {
	var input string
	for g.running {
		fmt.Scanln(&input)
		switch input {
		case "w":
			g.snake.changeDirection(Point{0, -1})
		case "a":
			g.snake.changeDirection(Point{-1, 0})
		case "s":
			g.snake.changeDirection(Point{0, 1})
		case "d":
			g.snake.changeDirection(Point{1, 0})
		case "q":
			g.running = false
		}
	}
}

func (g *Game) Update() {
	if !g.snake.alive {
		g.running = false
		return
	}

	g.snake.move()

	if g.snake.body[0] == g.food {
		g.snake.body = append([]Point{g.food}, g.snake.body...)
		g.food = generateFood()
		g.score++
	} else if len(g.snake.body) > 1 {
		g.snake.body = g.snake.body[:len(g.snake.body)-1]
	}

	g.snake.checkCollision()
}

func (s *Snake) move() {
	head := s.body[0]
	newHead := Point{head.x + s.dir.x, head.y + s.dir.y}
	s.body = append([]Point{newHead}, s.body[:len(s.body)-1]...)
}

func (s *Snake) checkCollision() {
	head := s.body[0]

	if head.x < 0 || head.x >= width || head.y < 0 || head.y >= height {
		s.alive = false
	}

	for _, p := range s.body[1:] {
		if p == head {
			s.alive = false
		}
	}
}

func (g *Game) Draw() {
	clearScreen()

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if g.isSnake(x, y) {
				fmt.Print("O")
			} else if g.food.x == x && g.food.y == y {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Printf("Score: %d\n", g.score)
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") // use "cmd" with "/c" and "cls" for Windows
	if err := cmd.Run(); err != nil {
		cmd = exec.Command("clear") // use "clear" for Unix-like systems
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (g *Game) isSnake(x, y int) bool {
	for _, p := range g.snake.body {
		if p.x == x && p.y == y {
			return true
		}
	}
	return false
}

func (s *Snake) changeDirection(dir Point) {
	// Prevent reversing direction
	if len(s.body) > 1 && s.body[0].x+dir.x == s.body[1].x && s.body[0].y+dir.y == s.body[1].y {
		return
	}
	s.dir = dir
}
