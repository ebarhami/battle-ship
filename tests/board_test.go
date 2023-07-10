package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"app/game"
)

var (
	board *game.Board
	ship1 *game.Ship
	ship2 *game.Ship
)

func initTest() {
	board = game.NewBoard(5)
	ship1 = game.NewShip(game.Coordinate{X: 1, Y: 1})
	ship2 = game.NewShip(game.Coordinate{X: 3, Y: 2})

	board.PlaceShip(ship2)
	board.PlaceShip(ship1)
}

func TestFireCoordinate_Valid(t *testing.T) {
	initTest()

	c1 := game.Coordinate{X: 1, Y: 1}
	c2 := game.Coordinate{X: 4, Y: 3}

	assert.Equal(t, len(board.Ships), 2)
	assert.Equal(t, 'B', board.Grid[c1.X][c1.Y]) // battleship is still alive
	valid := board.FireCoordinate(c1)
	assert.True(t, valid)                        // hit
	assert.Equal(t, 'X', board.Grid[c1.X][c1.Y]) // battleship is dead

	valid = board.FireCoordinate(c1)
	assert.False(t, valid) // idempotent, hit the same spot results nothing
	valid = board.FireCoordinate(c2)
	assert.False(t, valid) // miss

	hit := 0
	for _, ship := range board.Ships {
		if ship.IsHit {
			hit++
		}
	}
	assert.Equal(t, hit, 1)

	board.Print()
}

func TestFireCoordinate_InvalidCoordinate(t *testing.T) {
	initTest()

	c1 := game.Coordinate{X: -1, Y: 1}
	c2 := game.Coordinate{X: 5, Y: 1}
	valid := board.FireCoordinate(c1)
	assert.False(t, valid)
	valid = board.FireCoordinate(c2)
	assert.False(t, valid)

	board.Print()
}

func TestPlaceShip_Valid(t *testing.T) {
	initTest()

	// Place the third ship
	ship3 := game.NewShip(game.Coordinate{X: 4, Y: 0})
	success := board.PlaceShip(ship3)

	assert.True(t, success)
	assert.Equal(t, 3, len(board.Ships))
	assert.Equal(t, 3, len(board.GridToShips))
}

func TestPlaceShip_PlaceShipsInTheSameCoordinate(t *testing.T) {
	initTest()

	// Place the same ship2
	success := board.PlaceShip(ship2)

	assert.False(t, success)
	assert.Equal(t, 2, len(board.Ships))
	assert.Equal(t, 2, len(board.GridToShips))
}
