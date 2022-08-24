package main

import (
	"github.com/elvin-tacirzade/clean-architecture/pkg/app"
	"github.com/elvin-tacirzade/clean-architecture/pkg/config"
	"log"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}
	var a app.App
	a.Start()
}
