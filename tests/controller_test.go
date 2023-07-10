package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"app/game"
)

var (
	g          *game.Game
	s1, s2, s3 *game.Ship
	c1, c2, c3 game.Coordinate
	p1, p2, p3 *game.Player
)

func initTestGame() {
	g = game.NewGame(3, 5)
	p1, p2, p3 = g.Players[0], g.Players[1], g.Players[2]
	c1, c2, c3 = game.Coordinate{X: 1, Y: 1}, game.Coordinate{X: 2, Y: 4}, game.Coordinate{X: 4, Y: 3}
	s1, s2, s3 = game.NewShip(c1), game.NewShip(c2), game.NewShip(c3)
	p1.Board.PlaceShip(s1)
	p2.Board.PlaceShip(s2)
	p3.Board.PlaceShip(s3)
}

func TestGameWinnerP1(t *testing.T) {
	initTestGame()

	p1.Hit(p2, c2) // hit p2
	p1.Hit(p3, c3) // hit p3
	p2.Hit(p3, c1) // miss
	p3.Hit(p2, c1) // miss

	winner, isDraw := g.GetWinnerSoFar()

	assert.Equal(t, winner.Id, p1.Id)
	assert.False(t, isDraw)
}

func TestGameDraw(t *testing.T) {
	initTestGame()

	p1.Hit(p2, c2) // hit p2
	p2.Hit(p3, c3) // hit p3
	p3.Hit(p1, c1) // hit p1

	winner, isDraw := g.GetWinnerSoFar()

	assert.Nil(t, winner)
	assert.True(t, isDraw)
}
