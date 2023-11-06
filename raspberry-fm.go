package main

import (
	"flag"
	"fmt"

	"webimizer.dev/raspberry-fm/core"
)

func main() {
	fmt.Println("Welcome to Raspberry FM. Streaming your personal Web radio station using Raspberry PI.")

	port := flag.Int("port", 8080, "port number")

	flag.Parse()

	core.StartServer(*port)
}
