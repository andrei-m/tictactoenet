package tictactoenet

import (
	"errors"
	"math/bits"
)

type Game struct {
	// when false, it is playerA's turn
	playerBsTurn bool
	/*
	   Each state uint represents placements on the board:
	   1 2 4
	   8 16 32
	   64 128 256
	*/
	playerAState uint
	playerBState uint
}

func (g Game) playerAWins() bool {
	return isWin(g.playerAState)
}

func (g Game) playerBWins() bool {
	return isWin(g.playerBState)
}

func (g Game) isDraw() bool {
	return g.playerAState+g.playerBState == 511
}

var errInvalidMove = errors.New("invalid move")

func Move(game Game, move uint) (Game, error) {
	if !isValidMove(game, move) {
		return Game{}, errInvalidMove
	}
	if game.playerBsTurn {
		return Game{
			playerBsTurn: false,
			playerAState: game.playerAState,
			playerBState: game.playerBState | move,
		}, nil
	} else {
		return Game{
			playerBsTurn: true,
			playerAState: game.playerAState | move,
			playerBState: game.playerBState,
		}, nil
	}
}

func isValidMove(game Game, move uint) bool {
	if move > 256 || bits.OnesCount(move) != 1 {
		// must specify exactly one of the first 9 bits
		return false
	}
	if game.playerAState&move > 0 || game.playerBState&move > 0 {
		// space is taken
		return false
	}
	return true
}

var winningStates = []uint{
	// horizontal
	7, 56, 448,
	// vertical
	73, 146, 292,
	// diagonal
	84, 273,
}

func isWin(state uint) bool {
	for _, winningState := range winningStates {
		if state&winningState == winningState {
			return true
		}
	}
	return false
}
