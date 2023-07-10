package game

import (
	"fmt"
)

type Board struct {
	Grid        [][]rune
	Ships       []*Ship
	GridToShips map[Coordinate]*Ship
}

func NewBoard(size int) *Board {
	grid := make([][]rune, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]rune, size)
		for j := 0; j < size; j++ {
			grid[i][j] = '_'
		}
	}
	return &Board{
		Grid:        grid,
		Ships:       make([]*Ship, 0),
		GridToShips: make(map[Coordinate]*Ship),
	}
}

func (b *Board) PlaceShip(ship *Ship) bool {
	canPlace := b.canPlace(ship)

	// place the ship
	b.Ships = append(b.Ships, ship)
	b.GridToShips[ship.Position] = ship
	b.Grid[ship.Position.X][ship.Position.Y] = 'B'

	return canPlace
}

func (b *Board) FireCoordinate(coord Coordinate) (isShipHit bool) {
	isOutOfBound := !(coord.IsValid() && coord.X < len(b.Grid) && coord.Y < len(b.Grid))
	if isOutOfBound {
		return false
	}
	isAlreadyFired := b.Grid[coord.X][coord.Y] == 'O' || b.Grid[coord.X][coord.Y] == 'X'
	if isAlreadyFired {
		return false // idempotent, do nothing
	}

	if ship, ok := b.GridToShips[coord]; ok {
		b.Grid[coord.X][coord.Y] = 'X'
		ship.Hit()
		return true
	} else {
		b.Grid[coord.X][coord.Y] = 'O'
		return false
	}
}

func (b *Board) Print() {
	for i := 0; i < len(b.Grid); i++ {
		for j := 0; j < len(b.Grid); j++ {
			fmt.Printf("%c", b.Grid[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func (b *Board) canPlace(ship *Ship) bool {
	coord := ship.Position
	if b.GridToShips[coord] != nil { // Already have a ship in this coordinate
		return false
	}

	return true
}
