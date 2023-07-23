package api

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

func Start() {
	log.Println("hello")

	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
