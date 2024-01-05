package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	client := NewClient("", "")
	gui := NewGui(app, *client)
	gui.Run()

}
