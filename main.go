package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func main() {
	log.Println("hello")

	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
