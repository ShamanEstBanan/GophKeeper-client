package main

import (
	"GophKeeper-client/internal/app"
	"log"
	"os"
)

func main() {
	a := app.New()
	if err := a.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
