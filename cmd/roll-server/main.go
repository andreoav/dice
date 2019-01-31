package main

import (
	"log"

	"github.com/andreoav/dice/internal/app/rollserver"
)

func main() {
	config := &rollserver.Configuration{
		Address: "localhost:50051",
	}

	log.Fatalln(rollserver.Start(config))
}
