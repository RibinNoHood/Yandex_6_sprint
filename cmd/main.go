package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "Server: ", log.LstdFlags|log.Lshortfile)
	newServer := server.LaunchServer(logger)
	err := server.StartServer(newServer)
	if err != nil {
		logger.Fatal(err)
	}

}
