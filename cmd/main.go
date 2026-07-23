package main

import (
	"fmt"
	"log"

	"github.com/alexnakagama/redis-clone-go/internal/server"
)

func main() {
	port := ":6379"
	fmt.Println("Server running in port", port)

	s, err := server.NewServer(port)
	if err != nil {
		log.Println(err)
	}

	err = s.Start()
	if err != nil {
		log.Println(err)
	}
}
