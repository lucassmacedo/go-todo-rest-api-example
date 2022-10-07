package main

import (
	"github.com/lucassmacedo/go-todo-rest-api-example/app"
)

func main() {
	//config := config.GetConfig()

	app := &app.App{}
	//app.Initialize(config)
	app.Run(":3000")
}
