package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"app/game"
)

const (
	boardSize = 5
)

func TestPlayerHit_InvalidPlayer(t *testing.T) {
	player1 := game.NewPlayer(boardSize)
	// hit himself is not allowed
	err := player1.Hit(player1, game.NewCoordinate(1, 1))
	assert.NotNil(t, err)
}

func TestPlayerHit_OutOfBound(t *testing.T) {
	player1 := game.NewPlayer(boardSize)
	player2 := game.NewPlayer(boardSize)
	err := player1.Hit(player2, game.NewCoordinate(5, 1))
	assert.NotNil(t, err)
	err = player1.Hit(player2, game.NewCoordinate(3, 10))
	assert.NotNil(t, err)
	err = player1.Hit(player2, game.NewCoordinate(-1, 0))
	assert.NotNil(t, err)
}

func TestPlayerHit_Hit(t *testing.T) {
	player1 := game.NewPlayer(boardSize)
	player2 := game.NewPlayer(boardSize)
	player2.Board.PlaceShip(game.NewShip(game.NewCoordinate(2, 2)))

	// Miss hit, score is still zero
	err := player1.Hit(player2, game.NewCoordinate(3, 2))
	assert.Nil(t, err)
	assert.Equal(t, 0, player1.Score)

	// Hit the ship, score updated
	err = player1.Hit(player2, game.NewCoordinate(2, 2))
	assert.Nil(t, err)
	assert.Equal(t, 1, player1.Score)

	// hit the same spot, the score should remain the same
	err = player1.Hit(player2, game.NewCoordinate(2, 2))
	assert.Nil(t, err)
	assert.Equal(t, 1, player1.Score)
}
