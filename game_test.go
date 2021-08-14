package tictactoenet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Move(t *testing.T) {
	game := Game{}
	game, err := Move(game, 1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), game.playerAState)
	assert.False(t, game.playerAWins())
	assert.False(t, game.playerBWins())

	game, err = Move(game, 2)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), game.playerAState)
	assert.Equal(t, uint(2), game.playerBState)
	assert.False(t, game.playerAWins())
	assert.False(t, game.playerBWins())

	game, err = Move(game, 8)
	assert.NoError(t, err)
	assert.Equal(t, uint(9), game.playerAState)
	assert.Equal(t, uint(2), game.playerBState)
	assert.False(t, game.playerAWins())
	assert.False(t, game.playerBWins())

	game, err = Move(game, 64)
	assert.NoError(t, err)
	assert.Equal(t, uint(9), game.playerAState)
	assert.Equal(t, uint(66), game.playerBState)
	assert.False(t, game.playerAWins())
	assert.False(t, game.playerBWins())

	game, err = Move(game, 16)
	assert.NoError(t, err)
	assert.Equal(t, uint(25), game.playerAState)
	assert.Equal(t, uint(66), game.playerBState)
	assert.False(t, game.playerAWins())
	assert.False(t, game.playerBWins())

	game, err = Move(game, 128)
	assert.NoError(t, err)
	assert.Equal(t, uint(25), game.playerAState)
	assert.Equal(t, uint(194), game.playerBState)
	assert.False(t, game.playerAWins())
	assert.False(t, game.playerBWins())

	game, err = Move(game, 256)
	assert.NoError(t, err)
	assert.Equal(t, uint(281), game.playerAState)
	assert.Equal(t, uint(194), game.playerBState)
	assert.True(t, game.playerAWins())
	assert.False(t, game.playerBWins())
}

func Test_isValidMove(t *testing.T) {
	t.Run("out of range", func(t *testing.T) {
		assert.False(t, isValidMove(Game{}, 512))
	})

	t.Run("no squares", func(t *testing.T) {
		assert.False(t, isValidMove(Game{}, 0))
	})

	t.Run("multiple squares", func(t *testing.T) {
		assert.False(t, isValidMove(Game{}, 3))
	})

	t.Run("overlaps with player A", func(t *testing.T) {
		game := Game{
			playerAState: 36,
		}
		assert.False(t, isValidMove(game, 4))
	})

	t.Run("overlaps with player B", func(t *testing.T) {
		game := Game{
			playerAState: 36,
			playerBState: 9,
		}
		assert.False(t, isValidMove(game, 1))
	})

	t.Run("OK move", func(t *testing.T) {
		game := Game{
			playerAState: 36,
			playerBState: 9,
		}
		assert.True(t, isValidMove(game, 2))
	})
}
