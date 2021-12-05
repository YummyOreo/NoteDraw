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
	container1 := container.NewVBox(filler)

	var files []structs.NoteDrawFile

	btnMake := widgets.MakeButton(widgets.Button{Text: "Make Note", Func: func() {

		var name string

		MakeNoteWindow := a.NewWindow("Make Note")

		FileName := widget.NewEntry()

		FileName.Resize(fyne.NewSize(300, FileName.MinSize().Height))

		ButtonSubmit := widgets.MakeButton(widgets.Button{Text: "Make Note", Func: func() {
			name = FileName.Text
			MakeNoteWindow.Close()

			file := structs.NoteDrawFile{Name: name, LastModified: structs.Date{Month: 01, Day: 02, TimeHour: 03, TimeMin: 14}, Prev: "THis thisthsoentuhosenatuhsnoeat ntaouhesnteohusnt snaotuhoasen"}
			files = append(files, file)

			container1.Add(MakeFileCard(file))
			container1.Refresh()
		}})

		MakeNoteWindow.Resize(fyne.NewSize(float32(300), float32(100)))

		ContainerForm := container.NewVBox(FileName, ButtonSubmit)

		MakeNoteWindow.SetContent(ContainerForm)

		MakeNoteWindow.Show()
	}})

	text1 := widget.NewLabel("test")

	w.SetContent(container.NewVBox(btnMake, container.NewHBox(container1, text1)))

	w.SetCloseIntercept(func() {
		fmt.Println("closed")
		a.Quit()
	})

	w.ShowAndRun()
}

func MakeFileCard(file structs.NoteDrawFile) *fyne.Container {
	card1 := widget.NewCard(file.Name, fmt.Sprintf("%02d", file.LastModified.Month)+"/"+fmt.Sprintf("%02d", file.LastModified.Day)+" "+fmt.Sprintf("%02d", file.LastModified.TimeHour)+":"+fmt.Sprintf("%02d", file.LastModified.TimeMin), canvas.NewText(file.Prev, color.Gray{Y: 100}))
	btn1 := widgets.MakeButton(widgets.Button{Text: "Open File"})
	VBox1 := container.NewVBox(card1, btn1)
	return VBox1
}
