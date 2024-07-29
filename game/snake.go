package game

type Snake struct {
	Body  []Point
	Dir   Point
	Alive bool
}

func (s *Snake) Move() {
	head := s.Body[0]
	newHead := Point{head.X + s.Dir.X, head.Y + s.Dir.Y}
	s.Body = append([]Point{newHead}, s.Body...)
}

func (s *Snake) CheckCollision() {
	head := s.Body[0]

	if head.X < 0 || head.X >= width || head.Y < 0 || head.Y >= height {
		s.Alive = false
	}

	for _, p := range s.Body[1:] {
		if p == head {
			s.Alive = false
		}
	}
}

func (s *Snake) ChangeDirection(dir Point) {
	if len(s.Body) > 1 && s.Body[0].X+dir.X == s.Body[1].X && s.Body[0].Y+dir.Y == s.Body[1].Y {
		return
	}
	s.Dir = dir
}
