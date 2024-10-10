package main

import (
	"dune-imperium-service/internal/app"
)

func main() {

	application := &app.App{}
	application.Initialize()
	application.Run()
}
