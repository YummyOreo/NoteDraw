package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	//widgets "NoteDraw/mods"
)

func main() {
	a := app.New()

	w := a.NewWindow("My Window")

	with := 300
	height := 300

	multiH := 2
	multiW := 2

	w.Resize(fyne.NewSize(float32(with*multiW), float32(height*multiH)))

	card1 := widget.NewCard("Name Of file", "Last date motified", canvas.NewText("This is a prevew", color.Gray{Y: 100}))
	card2 := widget.NewCard("Name Of file", "Last date motified", canvas.NewText("This is a prevew", color.Gray{Y: 100}))
	card3 := widget.NewCard("Name Of file", "Last date motified", canvas.NewText("This is a prevew", color.Gray{Y: 100}))
	card4 := widget.NewCard("Name Of file", "Last date motified", canvas.NewText("This is a prevew", color.Gray{Y: 100}))
	card5 := widget.NewCard("Name Of file", "Last date motified", canvas.NewText("This is a prevew", color.Gray{Y: 100}))
	text1 := widget.NewLabel("test")

	w.SetContent(container.NewHBox(container.NewVBox(card1, card2, card3, card4, card5), text1))

	w.SetCloseIntercept(func() {
		fmt.Println("closed")
		w.Close()
	})

	w.ShowAndRun()

}
