package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"app/game"
)

func TestHitShip(t *testing.T) {
	ship := game.NewShip(game.Coordinate{})

	assert.False(t, ship.IsHit)
	ship.Hit()
	assert.True(t, ship.IsHit)
	
	// should be idempotent
	ship.Hit()
	assert.True(t, ship.IsHit)
}
