package main

import (
	"fmt"
	"log"

	"github.com/alexnakagama/redis-clone-go/internal/server"
)

func main() {
	fmt.Println("Welcome to the TCP Server")

	s, err := server.NewServer(":6379")
	if err != nil {
		log.Println(err)
	}

	err = s.Start()
	if err != nil {
		log.Println(err)
	}
}
