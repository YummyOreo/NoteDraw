package main

import (
	"fmt"
	"image/color"

	note "NoteDraw/Note"
	"NoteDraw/exporting"
	"NoteDraw/loading"
	"NoteDraw/saving"
	"NoteDraw/snippets"
	"NoteDraw/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// makes the app
	a := app.New()

	// makes the main windew
	w := a.NewWindow("NoteDraw")

	// centers the window and resizes it to most of the screen
	w.CenterOnScreen()
	w.Resize(snippets.WindowSize(0.8))

	// makes the container for the top
	ContainerHead := container.NewVBox()

	// makes the container for files
	ContainerFiles := container.NewVBox()

	// makes the container for the main content
	ContainerShowContent := container.NewMax()

	// tracking the last content (for the main content)
	LastContainer := new(structs.LastContent)

	// tracking the files
	files := new(structs.Files)
	files.Files = make(map[string]structs.NoteDrawFile)
	files.Cards = make(map[string]*widget.Card)

	// tracking the current file
	current := new(structs.CurrentFile)

	// the text on the top of the window
	Header := widget.NewLabel("NoteDraw")
	Header.Alignment = fyne.TextAlignCenter

	// makes the main button
	ButtonMakeNote := note.MakeButton(a, files, current, LastContainer, ContainerFiles, ContainerShowContent)

	// makes a shortcut for saving the current note
	w.Canvas().AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: desktop.ControlModifier}, func(shortcut fyne.Shortcut) {
		saving.Save(files, current)
	})

	// exporting (not done)
	w.Canvas().AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyE, Modifier: desktop.ControlModifier}, func(shortcut fyne.Shortcut) {
		exporting.Shortcut(files, current)
	})

	// sets the button that makes the notes to be meduim importance
	ButtonMakeNote.Importance = widget.MediumImportance

	// makes it so you can scroll through the files
	ContainerScroll := container.NewVScroll(ContainerFiles)

	// adds the header and the button to make the notes
	ContainerHead.Add(Header)

	ContainerHead.Add(canvas.NewLine(color.RGBA{R: 24, G: 24, B: 24, A: 190}))

	ContainerHead.Add(container.NewHBox(ButtonMakeNote, loading.LoadFile(w, files, current, LastContainer, ContainerFiles, ContainerShowContent)))

	ContainerHead.Add(canvas.NewLine(color.RGBA{R: 24, G: 24, B: 24, A: 190}))

	// makes the main container, and makes it so they are formated correctly
	ContainerMain := container.NewBorder(ContainerHead, nil, ContainerScroll, nil, ContainerShowContent)

	// sets the main windows content to the main container
	w.SetContent(ContainerMain)

	w.SetCloseIntercept(func() {
		fmt.Println("closed")
		a.Quit()
	})

	w.ShowAndRun()
}
