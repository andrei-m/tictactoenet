package main

import (
	"log"
	"os"
	"strconv"

	"github.com/andrei-m/tictactoenet"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("failed to parse port: %v", err)
	}

	if err := tictactoenet.Listen(port); err != nil {
		log.Fatalf("%v", err)
	}
}
