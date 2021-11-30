package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"

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

	w.SetContent(widget.MakeIcon(theme.AccountIcon()))

	w.SetCloseIntercept(func() {
		fmt.Println("closed")
		w.Close()
	})

	w.ShowAndRun()

}
