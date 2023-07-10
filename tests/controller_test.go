package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"app/game"
)

var (
	g      *game.Game
	s1, s2 *game.Ship
	c1, c2 game.Coordinate
	p1, p2 *game.Player
)

func initTestGame() {
	g = game.NewGame(2, 5)
	p1, p2 = g.Players[0], g.Players[1]
	c1, c2 = game.Coordinate{X: 1, Y: 1}, game.Coordinate{X: 2, Y: 4}
	s1, s2 = game.NewShip(c1), game.NewShip(c2)
	p1.Board.PlaceShip(s1)
	p2.Board.PlaceShip(s2)
}

func TestGameWinnerP1(t *testing.T) {
	initTestGame()

	p1.Hit(p2, c2) // hit p2

	winner, isDraw := g.GetWinnerSoFar()

	assert.Equal(t, 1, p1.Score)
	assert.Equal(t, winner.Id, p1.Id)
	assert.False(t, isDraw)
}

func TestGameDraw(t *testing.T) {
	initTestGame()

	p1.Hit(p2, c2) // hit p2
	p2.Hit(p1, c1) // hit p3

	winner, isDraw := g.GetWinnerSoFar()

	assert.Nil(t, winner)
	assert.True(t, isDraw)
}
