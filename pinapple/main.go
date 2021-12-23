package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"time"

	note "NoteDraw/Note"
	"NoteDraw/snippets"
	"NoteDraw/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
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
	Header := widget.NewLabel("NoteDraw")
	Header.Alignment = fyne.TextAlignCenter

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
		result, err := cfdutil.ShowSaveFileDialog(cfd.DialogConfig{
			Title: "Save A File",
			Role:  "SaveFileExample",
			FileFilters: []cfd.FileFilter{
				{
					DisplayName: "JSON Files (*.json)",
					Pattern:     "*.json",
				},
				{
					DisplayName: "NoteDraw Files (*.nd)",
					Pattern:     "*.nd",
				},
			},
			SelectedFileFilterIndex: 1,
			FileName:                current.FileName + ".nd",
			DefaultExtension:        "nd",
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		if result == "" {
			return
		}

		TempFile := structs.SaveFile{Name: current.FileName, LastModified: files.Files[current.FileName].LastModified}
		for _, v := range files.Files[current.FileName].Content {
			switch v.Type {
			case "title":
				TempFile.Content = append(TempFile.Content, structs.SaveType{Type: v.Type, Data: v.Title.Text.Text})
			case "paragraph":
				TempFile.Content = append(TempFile.Content, structs.SaveType{Type: v.Type, Data: v.Paragraph.Text.Text})
			case "draw":
				lines := []structs.SaveDraw{}
				for _, v := range v.Drawing.Canvas.Line {
					lines = append(lines, structs.SaveDraw{Pos1X: int(v.Position1.X), Pos1Y: int(v.Position1.Y), Pos2X: int(v.Position2.X), Pos2Y: int(v.Position2.Y)})
				}
				TempFile.Content = append(TempFile.Content, structs.SaveType{Type: v.Type, Lines: lines})
			}
		}
		b, err := json.Marshal(TempFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
		ioutil.WriteFile(result, b, os.ModePerm)
	})

	// importing (not done)
	ImportNote := snippets.MakeButton(snippets.Button{Text: "Import", Func: func() {
		result, err := cfdutil.ShowOpenFileDialog(cfd.DialogConfig{
			Title: "Open A File",
			Role:  "OpenFileExample",
			FileFilters: []cfd.FileFilter{
				{
					DisplayName: "JSON Files (*.json)",
					Pattern:     "*.json",
				},
				{
					DisplayName: "NoteDraw Files (*.nd)",
					Pattern:     "*.nd",
				},
			},
			SelectedFileFilterIndex: 2,
			FileName:                current.FileName + ".nd",
			DefaultExtension:        "nd",
		})
		if err != nil {
			return
		}

		// Open our jsonFile
		jsonFile, err := os.Open(result)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Println(err)
			return
		}

		if jsonFile.Name() == "" {
			return
		}

		byteValue, _ := ioutil.ReadAll(jsonFile)
		tempStruct := structs.SaveFileJson{}
		json.Unmarshal(byteValue, &tempStruct)

		file := structs.NoteDrawFile{Name: tempStruct.Name, LastModified: tempStruct.LastModified}

		for _, v := range files.Files {
			if v.Name == tempStruct.Name {
				// if it is, make a popup saying that it is already taken
				popup := widget.NewModalPopUp(
					widget.NewLabel("That already exists"),
					w.Canvas(),
				)
				popup.Show()
				// wait 2 seconds
				time.Sleep(2 * time.Second)
				// hide the popup, and stop the code for doing anything else
				popup.Hide()
				return
			}
		}

		for _, v := range tempStruct.Content {
			switch v.Type {
			case "paragraph":
				Ph := widget.NewMultiLineEntry()
				Ph.Wrapping = fyne.TextWrapBreak
				Ph.Text = v.Data
				file.Content = append(file.Content, structs.NoteType{Type: v.Type, Paragraph: structs.Paragraph{Text: Ph}})
			case "title":
				// makes a new multily line text entry
				Title := widget.NewEntry()
				Title.TextStyle.Bold = true
				Title.Text = v.Data

				// makes a new type, and sets the text to be the new multiline entry, then appends it to the types
				file.Content = append(file.Content, structs.NoteType{Type: "title", Title: structs.Title{Text: Title}})
			case "draw":
				drawing := structs.Drawing{}
				lines := new(structs.LineList)
				for _, v := range v.Lines {
					line := canvas.NewLine(color.Black)
					line.StrokeWidth = 2

					line.Position1.X = float32(v.Pos1X)
					line.Position1.Y = float32(v.Pos1Y)
					line.Position2.X = float32(v.Pos2X)
					line.Position2.Y = float32(v.Pos2Y)
					lines.Line = append(lines.Line, line)
				}
				drawing.Canvas = lines
				file.Content = append(file.Content, structs.NoteType{Type: v.Type, Drawing: drawing})
			}
		}

		// defer the closing of our jsonFile so that we can parse it later on
		defer jsonFile.Close()

		// save the file to a map
		files.Files[file.Name] = file

		// make the card and append it to the container
		ContainerFiles.Add(note.MakeFileCard(file, files, current, ContainerShowContent, LastContainer))
		ContainerFiles.Objects = snippets.MoveInt(ContainerFiles.Objects, len(ContainerFiles.Objects)-1, 0)
		ContainerFiles.Refresh()
		return
	}})

	// sets the button that makes the notes to be meduim importance
	ButtonMakeNote.Importance = widget.MediumImportance

	// makes it so you can scroll through the files
	ContainerScroll := container.NewVScroll(ContainerFiles)

	// adds the header and the button to make the notes
	ContainerHead.Add(Header)

	ContainerHead.Add(canvas.NewLine(color.RGBA{R: 24, G: 24, B: 24, A: 190}))

	ContainerHead.Add(container.NewHBox(ButtonMakeNote, ImportNote))

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
