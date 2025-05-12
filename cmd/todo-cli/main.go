package main

import (
	"github.com/furkankorkmaz309/todo-cli/internal"
)

func main() {

	app, err := internal.NewApp()
	if err != nil {
		app.Logger.ErrorLog.Fatal("Failed to initialize app")
	}

	internal.Run(app)
}
