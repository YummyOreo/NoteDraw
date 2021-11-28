package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()

	w := a.NewWindow("My Window")

	with := 200
	height := 400

	multiH := 2
	multiW := 3

	w.Resize(fyne.NewSize(float32(with*multiW), float32(height*multiH)))
	w.ShowAndRun()

}
