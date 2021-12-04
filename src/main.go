package main

import (
	"fmt"
	"image/color"

	widgets "NoteDraw/mods"
	structs "NoteDraw/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("My Window")

	with := 300
	height := 300

	multiH := 2
	multiW := 2

	w.Resize(fyne.NewSize(float32(with*multiW), float32(height*multiH)))
	filler := widget.NewLabel("")
	contaiter1 := container.NewVBox(filler)

	var files []structs.NoteDrawFile

	btnMake := widgets.MakeButton(widgets.Button{Text: "Make Note", Func: func() {
		file := structs.NoteDrawFile{Name: "name", LastModified: structs.Date{Month: 1, Day: 2, TimeHour: 3, TimeMin: 4}, Prev: "THis thisthsoentuhosenatuhsnoeat ntaouhesnteohusnt snaotuhoasen"}
		files = append(files, file)

		contaiter1.Add(MakeFileCard(file))
		contaiter1.Refresh()
	}})

	text1 := widget.NewLabel("test")

	w.SetContent(container.NewVBox(btnMake, container.NewHBox(contaiter1, text1)))

	w.SetCloseIntercept(func() {
		fmt.Println("closed")
		w.Close()
	})

	w.ShowAndRun()
}

func MakeFileCard(file structs.NoteDrawFile) *fyne.Container {
	card1 := widget.NewCard(file.Name, file.Name, canvas.NewText(file.Prev, color.Gray{Y: 100}))
	btn1 := widgets.MakeButton(widgets.Button{Text: "Open File"})
	VBox1 := container.NewVBox(card1, btn1)
	return VBox1
}
