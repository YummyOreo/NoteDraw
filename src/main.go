package main

import (
	"fmt"
	"image/color"
	"time"

	snippets "NoteDraw/snippets"
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
	a := app.New()

	w := a.NewWindow("NoteDraw")

	w.CenterOnScreen()
	w.Resize(windowSize(0.8))

	ContainerMakeNote := container.NewVBox()

	ContainerFiles := container.NewVBox()

	ContainerShowContent := container.NewMax()

	LastContainer := new(structs.LastContent)

	files := new(structs.Files)
	files.Files = make(map[string]structs.NoteDrawFile)
	files.Cards = make(map[string]*widget.Card)

	current := new(structs.CurrentFile)

	Header := canvas.NewText("NoteDraw", color.White)
	Header.Alignment = fyne.TextAlignCenter
	Header.TextSize = 20

	ButtonMakeNote := MakeButton(a, files, ContainerFiles, ContainerShowContent, LastContainer, current)

	w.Canvas().AddShortcut(&desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: desktop.ControlModifier}, func(shortcut fyne.Shortcut) {
		TempFile := structs.NoteDrawFile{}
		TempFile.Content = current.Text.Text
		TempFile.Name = files.Files[current.FileName].Name
		TempFile.LastModified = structs.Date{Month: int(time.Now().Month()), Day: time.Now().Day(), TimeHour: time.Now().Hour(), TimeMin: time.Now().Minute()}
		TempFile.Prev = snippets.GetPrev(TempFile)
		if TempFile.LastModified.TimeHour > 12 {
			TempFile.LastModified.TimeHour = TempFile.LastModified.TimeHour - 12
		}
		files.Files[TempFile.Name] = TempFile
		UpdateCard(files.Files[TempFile.Name], files)
		fmt.Println(current.Text.Text)
	})

	ButtonMakeNote.Importance = widget.MediumImportance

	ContainerScroll := container.NewVScroll(ContainerFiles)

	ContainerMakeNote.Add(Header)
	ContainerMakeNote.Add(ButtonMakeNote)

	ContainerMain := container.NewBorder(ContainerMakeNote, nil, ContainerScroll, nil, ContainerShowContent)

	w.SetContent(ContainerMain)

	w.SetCloseIntercept(func() {
		fmt.Println("closed")
		a.Quit()
	})

	w.ShowAndRun()
}

func MakeButton(a fyne.App, files *structs.Files, ContainerFiles *fyne.Container, ContainerShowContent *fyne.Container, LastContainer *structs.LastContent, current *structs.CurrentFile) *widget.Button {
	ButtonMakeNote := snippets.MakeButton(snippets.Button{Text: "Make Note", Func: func() {

		var name string

		MakeNoteWindow := a.NewWindow("Make Note")

		FileName := widget.NewEntry()

		FileName.Resize(fyne.NewSize(300, FileName.MinSize().Height))

		ButtonSubmit := snippets.MakeButton(snippets.Button{Text: "Make Note", Func: func() {
			if FileName.Text == "" {
				return
			}
			name = FileName.Text
			MakeNoteWindow.Close()

			file := structs.NoteDrawFile{Name: name, LastModified: structs.Date{Month: int(time.Now().Month()), Day: time.Now().Day(), TimeHour: time.Now().Hour(), TimeMin: time.Now().Minute()}, Prev: "This is a preview"}
			if file.LastModified.TimeHour > 12 {
				file.LastModified.TimeHour = file.LastModified.TimeHour - 12
			}
			file.Prev = snippets.GetPrev(file)
			files.Files[name] = file
			ContainerFiles.Add(MakeFileCard(file, ContainerShowContent, LastContainer, files, current))
			ContainerFiles.Refresh()
			return
		}})

		MakeNoteWindow.Resize(fyne.NewSize(float32(300), float32(100)))

		ContainerForm := container.NewVBox(FileName, ButtonSubmit)

		MakeNoteWindow.SetContent(ContainerForm)

		MakeNoteWindow.Show()
		return
	}})
	return ButtonMakeNote
}

func MakeFileCard(file structs.NoteDrawFile, ContainerShowContent *fyne.Container, LastContainer *structs.LastContent, files *structs.Files, current *structs.CurrentFile) *fyne.Container {
	text := canvas.NewText(file.Prev, color.Gray{Y: 100})
	text.TextSize = 10
	card1 := widget.NewCard(files.Files[file.Name].Name, fmt.Sprintf("%02d", files.Files[file.Name].LastModified.Month)+"/"+fmt.Sprintf("%02d", files.Files[file.Name].LastModified.Day)+" "+fmt.Sprintf("%02d", file.LastModified.TimeHour)+":"+fmt.Sprintf("%02d", file.LastModified.TimeMin), text)
	files.Cards[file.Name] = card1
	btn1 := snippets.MakeButton(snippets.Button{Text: "Open File"})
	btn1.OnTapped = func() {
		ContainerShowContent.Remove(LastContainer.Content)
		ContainerShowContent.Refresh()
		ContainerShowContent.Add(MakeContent(file, LastContainer, current, files))
		ContainerShowContent.Refresh()
		return
	}
	VBox1 := container.NewVBox(card1, btn1)
	return container.NewWithoutLayout(VBox1)
}

func MakeContent(file structs.NoteDrawFile, LastContainer *structs.LastContent, current *structs.CurrentFile, files *structs.Files) *widget.Entry {
	text := widget.NewMultiLineEntry()
	text.Wrapping = fyne.TextWrapBreak
	text.Text = files.Files[file.Name].Content
	LastContainer.Content = text
	current.FileName = file.Name
	current.Text = text
	return text
}

func UpdateCard(file structs.NoteDrawFile, files *structs.Files) {
	text := canvas.NewText(files.Files[file.Name].Prev, color.Gray{Y: 100})
	text.TextSize = 10
	files.Cards[file.Name].SetContent(text)
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
