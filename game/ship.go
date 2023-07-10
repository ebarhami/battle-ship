package game

type Coordinate struct {
	X int
	Y int
}

func (c *Coordinate) IsValid() bool {
	return c.X >= 0 && c.Y >= 0
}

type Ship struct {
	Name     string
	IsHit    bool
	Position Coordinate
}

func NewShip(name string, position Coordinate) *Ship {
	return &Ship{
		Name:     name,
		Position: position,
	}
}

func (s *Ship) Hit() {
	s.IsHit = true
}
