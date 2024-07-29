package game

import (
	"fmt"
	"gosnake/constants"
)

type Snake struct {
	Body  []Point
	Dir   Point
	Alive bool
}

func (s *Snake) Move() {
	head := s.Body[0]
	newHead := Point{head.X + s.Dir.X, head.Y + s.Dir.Y}
	s.Body = append([]Point{newHead}, s.Body...) // Add new head to the body
	fmt.Printf("Moved to new head position: %+v\n", newHead)
}

func (s *Snake) CheckCollision() {
	head := s.Body[0]

	// Check boundary collision
	if head.X < 0 || head.X >= constants.Width || head.Y < 0 || head.Y >= constants.Height {
		s.Alive = false
		fmt.Println("Collision with boundary detected!")
		return
	}

	// Check self-collision
	for _, p := range s.Body[1:] {
		if p == head {
			s.Alive = false
			fmt.Println("Collision with self detected!")
			return
		}
	}
}

func (s *Snake) ChangeDirection(dir Point) {
	if len(s.Body) > 1 && s.Body[0].X+dir.X == s.Body[1].X && s.Body[0].Y+dir.Y == s.Body[1].Y {
		fmt.Println("Invalid direction change: reversing direction")
		return
	}
	s.Dir = dir
	fmt.Printf("Direction changed to: %+v\n", dir)
}
