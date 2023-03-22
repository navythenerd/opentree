package main

import (
	"fmt"
	"log"
	"net/http"
	"opentree/app"
	"os"
)

func main() {
	fmt.Println("OpenTree")

	args := os.Args[1:]

	var config app.Config

	if len(args) == 1 {
		app.ReadJSON(&config, args[0])
	} else {
		app.ReadJSON(&config, "config.json")
	}

	opentree := app.New(config)

	fmt.Printf("listening on %s:%d\n", config.Address, config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", config.Address, config.Port), opentree))
}
