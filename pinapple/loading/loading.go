package loading

import (
	"NoteDraw/Note"
	"NoteDraw/snippets"
	"NoteDraw/structs"
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
)

func LoadFile(w fyne.Window, files *structs.Files, current *structs.CurrentFile, LastContainer *structs.LastContent, ContainerFiles *fyne.Container, ContainerShowContent *fyne.Container) *widget.Button {
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
		ContainerFiles.Add(Note.MakeFileCard(file, files, current, ContainerShowContent, LastContainer))
		ContainerFiles.Objects = snippets.Move(ContainerFiles.Objects, len(ContainerFiles.Objects)-1, 0)
		ContainerFiles.Refresh()
		return
	}})
	return ImportNote
}
