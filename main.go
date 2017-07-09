package main

import (
	"log"
	"os"

	"github.com/matmerr/litturl/server"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("web directory argument required")
		return
	}
	clientdir := os.Args[1]
	server.Start(clientdir)
}
