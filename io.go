package tictactoenet

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

func Listen(port int) error {
	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return errors.WithStack(err)
	}
	log.Printf("listening on port %d", port)
	defer l.Close()

	gameNo := 1
	for {
		conn, err := l.Accept()
		if err != nil {
			return errors.WithStack(err)
		}
		go handleConn(gameNo, conn)
		gameNo++
	}
}

func handleConn(gameNo int, conn net.Conn) {
	defer conn.Close()

	log.Printf("beginning game #%d", gameNo)
	game := Game{}
	for {
		if game.playerOsTurn {
			game = remoteMove(gameNo, game, conn)
		} else {
			game = localMove(gameNo, game)
		}

		if game.playerXWins() {
			fmt.Println("You win!")
			conn.Write([]byte("You lose!"))
			break
		} else if game.playerOWins() {
			fmt.Println("You lose!")
			conn.Write([]byte("You win!"))
			break
		} else if game.isDraw() {
			fmt.Println("It's a draw.")
			conn.Write([]byte("It's a draw."))
			break
		}
	}
}

func localMove(gameNo int, game Game) Game {
	return collectMove(gameNo, game, os.Stdin, os.Stderr)
}

func remoteMove(gameNo int, game Game, conn net.Conn) Game {
	log.Println("awaiting remote move...")
	return collectMove(gameNo, game, conn, conn)
}

func collectMove(gameNo int, game Game, in io.Reader, out io.Writer) Game {
	turn := "X"
	if game.playerOsTurn {
		turn = "O"
	}
	scanner := bufio.NewScanner(in)
	for {
		io.WriteString(out, game.String())
		io.WriteString(out, fmt.Sprintf("Game #%d move (player %s)> ", gameNo, turn))
		if !scanner.Scan() {
			continue
		}
		move, err := strconv.Atoi(scanner.Text())
		if err != nil {
			io.WriteString(out, "invalid input\n")
			continue
		}
		moved, err := Move(game, uint(move))
		if err != nil {
			io.WriteString(out, "invalid move\n")
			continue
		}
		return moved
	}
}
