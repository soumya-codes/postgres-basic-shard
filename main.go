package main

import (
	"log"

	"github.com/soumya-codes/postgres-static-shard/internal/router"
)

func main() {
	r := router.SetupRouter()
	log.Fatal(r.Run(":8080"))
}
