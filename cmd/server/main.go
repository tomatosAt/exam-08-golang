package main

import (
	"fmt"
	"log"
	"os"
)

var Version string

func main() {
	fmt.Println("Junior API is running ")
	if len(os.Args) < 2 {
		log.Fatal("Please specify a command: serve | migrate")
	}
	switch os.Args[1] {
	case "serve", "start":
		startServer()
	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
	}
}
