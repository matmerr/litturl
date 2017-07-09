package main

import (
	"os"
	"sync"

	"github.com/matmerr/litturl/server"
)

func main() {

	clientdir := os.Args[1]

	server.Start(clientdir)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
	}()
	wg.Wait()
}
