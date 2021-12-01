package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	widget "NotDraw/mods"
)

func main() {
	a := app.New()

	w := a.NewWindow("My Window")

	with := 300
	height := 300

	multiH := 2
	multiW := 2

	w.Resize(fyne.NewSize(float32(with*multiW), float32(height*multiH)))

	w.SetContent(widget.MakeLine(widget.Line{Color: color.Black}))

	w.SetCloseIntercept(func() {
		fmt.Println("closed")
		w.Close()
	})

	w.ShowAndRun()

}
