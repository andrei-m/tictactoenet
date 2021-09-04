package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/andrei-m/tictactoenet"
)

func main() {
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("failed to parse port: %v", err)
	}

	rand.Seed(time.Now().UnixNano())
	if err := tictactoenet.Listen(port); err != nil {
		log.Fatalf("%v", err)
	}
}
