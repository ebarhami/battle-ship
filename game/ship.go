package game

type Coordinate struct {
	X int
	Y int
}

func NewCoordinate(x, y int) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
	}
}

func (c *Coordinate) IsValid() bool {
	return c.X >= 0 && c.Y >= 0
}

type Ship struct {
	IsHit    bool
	Position Coordinate
}

func NewShip(position Coordinate) *Ship {
	return &Ship{
		Position: position,
	}
}

func (s *Ship) Hit() {
	s.IsHit = true
}
