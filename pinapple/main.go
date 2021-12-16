package main

import (
	"fmt"
	"image/color"
	"time"

	note "NoteDraw/Note"
	"NoteDraw/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/kbinani/screenshot"
)

func main() {
	// makes the app
	a := app.New()

	// makes the main windew
	w := a.NewWindow("NoteDraw")

	// centers the window and resizes it to most of the screen
	w.CenterOnScreen()
	w.Resize(windowSize(0.8))

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
	Header := canvas.NewText("NoteDraw", color.White)
	Header.Alignment = fyne.TextAlignCenter
	Header.TextSize = 20

	// makes the main button
	ButtonMakeNote := note.MakeButton(a, files, current, LastContainer, ContainerFiles, ContainerShowContent)

	// makes a shortcut for saving the current note
	w.Canvas().AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: desktop.ControlModifier}, func(shortcut fyne.Shortcut) {
		// makes a new temp note, that is blank
		TempFile := structs.NoteDrawFile{}

		// sets the types to the same of the current one
		TempFile.Content = current.Types

		// sets the name to the current file name
		TempFile.Name = files.Files[current.FileName].Name

		// sets the last modified to the current time
		TempFile.LastModified = structs.Date{Month: int(time.Now().Month()), Day: time.Now().Day(), TimeHour: time.Now().Hour(), TimeMin: time.Now().Minute()}

		// changes the time to 12 hour format
		if TempFile.LastModified.TimeHour > 12 {
			TempFile.LastModified.TimeHour = TempFile.LastModified.TimeHour - 12
		}

		// updates the map to be the updated one
		files.Files[TempFile.Name] = TempFile

		// updates the card
		UpdateCard(files.Files[TempFile.Name], files)
	})

	// exporting (not done)
	w.Canvas().AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyE, Modifier: desktop.ControlModifier}, func(shortcut fyne.Shortcut) {
		TempFile := files.Files[current.FileName]
		TempFile.LastModified = structs.Date{Month: int(time.Now().Month()), Day: time.Now().Day(), TimeHour: time.Now().Hour(), TimeMin: time.Now().Minute()}
		if TempFile.LastModified.TimeHour > 12 {
			TempFile.LastModified.TimeHour = TempFile.LastModified.TimeHour - 12
		}
	})

	// sets the button that makes the notes to be meduim importance
	ButtonMakeNote.Importance = widget.MediumImportance

	// makes it so you can scroll through the files
	ContainerScroll := container.NewVScroll(ContainerFiles)

	// adds the header and the button to make the notes
	ContainerHead.Add(Header)
	ContainerHead.Add(ButtonMakeNote)

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

func UpdateCard(file structs.NoteDrawFile, files *structs.Files) {
	files.Cards[file.Name].Subtitle = fmt.Sprintf("%02d", files.Files[file.Name].LastModified.Month) + "/" + fmt.Sprintf("%02d", files.Files[file.Name].LastModified.Day) + " " + fmt.Sprintf("%02d", file.LastModified.TimeHour) + ":" + fmt.Sprintf("%02d", file.LastModified.TimeMin)
	files.Cards[file.Name].Refresh()
}

func windowSize(part float32) fyne.Size {
	if screenshot.NumActiveDisplays() > 0 {
		// #0 is the main monitor
		bounds := screenshot.GetDisplayBounds(0)
		return fyne.NewSize(float32(bounds.Dx())*part, float32(bounds.Dy())*part)
	}
	return fyne.NewSize(800, 800)
}