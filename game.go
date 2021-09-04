package tictactoenet

import (
	"errors"
	"fmt"
	"math/bits"
)

type Game struct {
	// when false, it is playerX's turn
	playerOsTurn bool
	/*
	   Each state uint represents placements on the board:
	   1 2 4
	   8 16 32
	   64 128 256
	*/
	playerXState uint
	playerOState uint
}

func (g Game) playerXWins() bool {
	return isWin(g.playerXState)
}

func (g Game) playerOWins() bool {
	return isWin(g.playerOState)
}

func (g Game) isDraw() bool {
	return g.playerXState+g.playerOState == 511
}

func (g Game) String() string {
	state := make([]interface{}, 9)
	for i := range state {
		bit := 1 << i
		uintBit := uint(bit)
		if g.playerXState&uintBit == uintBit {
			state[i] = " X "
		} else if g.playerOState&uintBit == uintBit {
			state[i] = " O "
		} else {
			state[i] = fmt.Sprintf("%3d", bit)
		}
	}

	return fmt.Sprintf(`
%s|%s|%s
-------------
%s|%s|%s
-------------
%s|%s|%s
`, state...)
}

var errInvalidMove = errors.New("invalid move")

func Move(game Game, move uint) (Game, error) {
	if !isValidMove(game, move) {
		return Game{}, errInvalidMove
	}
	if game.playerOsTurn {
		return Game{
			playerOsTurn: false,
			playerXState: game.playerXState,
			playerOState: game.playerOState | move,
		}, nil
	} else {
		return Game{
			playerOsTurn: true,
			playerXState: game.playerXState | move,
			playerOState: game.playerOState,
		}, nil
	}
}

func isValidMove(game Game, move uint) bool {
	if move > 256 || bits.OnesCount(move) != 1 {
		// must specify exactly one of the first 9 bits
		return false
	}
	if game.playerXState&move > 0 || game.playerOState&move > 0 {
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
