package tictactoenet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Move(t *testing.T) {
	game := Game{}
	game, err := Move(game, 1)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), game.playerXState)
	assert.False(t, game.playerXWins())
	assert.False(t, game.playerOWins())

	game, err = Move(game, 2)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), game.playerXState)
	assert.Equal(t, uint(2), game.playerOState)
	assert.False(t, game.playerXWins())
	assert.False(t, game.playerOWins())

	game, err = Move(game, 8)
	assert.NoError(t, err)
	assert.Equal(t, uint(9), game.playerXState)
	assert.Equal(t, uint(2), game.playerOState)
	assert.False(t, game.playerXWins())
	assert.False(t, game.playerOWins())

	game, err = Move(game, 64)
	assert.NoError(t, err)
	assert.Equal(t, uint(9), game.playerXState)
	assert.Equal(t, uint(66), game.playerOState)
	assert.False(t, game.playerXWins())
	assert.False(t, game.playerOWins())

	game, err = Move(game, 16)
	assert.NoError(t, err)
	assert.Equal(t, uint(25), game.playerXState)
	assert.Equal(t, uint(66), game.playerOState)
	assert.False(t, game.playerXWins())
	assert.False(t, game.playerOWins())

	game, err = Move(game, 128)
	assert.NoError(t, err)
	assert.Equal(t, uint(25), game.playerXState)
	assert.Equal(t, uint(194), game.playerOState)
	assert.False(t, game.playerXWins())
	assert.False(t, game.playerOWins())

	game, err = Move(game, 256)
	assert.NoError(t, err)
	assert.Equal(t, uint(281), game.playerXState)
	assert.Equal(t, uint(194), game.playerOState)
	assert.True(t, game.playerXWins())
	assert.False(t, game.playerOWins())
}

func Test_Game_isDraw(t *testing.T) {
	game := Game{}
	game, err := Move(game, 1)
	assert.NoError(t, err)
	game, err = Move(game, 2)
	assert.NoError(t, err)
	game, err = Move(game, 8)
	assert.NoError(t, err)
	game, err = Move(game, 64)
	assert.NoError(t, err)
	game, err = Move(game, 32)
	assert.NoError(t, err)
	game, err = Move(game, 256)
	assert.NoError(t, err)
	game, err = Move(game, 4)
	assert.NoError(t, err)
	game, err = Move(game, 128)
	assert.NoError(t, err)
	game, err = Move(game, 16)
	assert.NoError(t, err)
	assert.True(t, game.isDraw())
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
			playerXState: 36,
		}
		assert.False(t, isValidMove(game, 4))
	})

	t.Run("overlaps with player B", func(t *testing.T) {
		game := Game{
			playerXState: 36,
			playerOState: 9,
		}
		assert.False(t, isValidMove(game, 1))
	})

	t.Run("OK move", func(t *testing.T) {
		game := Game{
			playerXState: 36,
			playerOState: 9,
		}
		assert.True(t, isValidMove(game, 2))
	})
}
