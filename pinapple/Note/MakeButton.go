package Note

import (
	"NoteDraw/snippets"
	"NoteDraw/structs"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MakeButton(a fyne.App, files *structs.Files, current *structs.CurrentFile, LastContainer *structs.LastContent, ContainerFiles *fyne.Container, ContainerShowContent *fyne.Container) *widget.Button {
	// makes the button for making a new Note
	ButtonMakeNote := snippets.MakeButton(snippets.Button{Text: "Make Note", Func: func() {

		// makes a var to store the name
		var name string

		// makes a new window (a pop up)
		MakeNoteWindow := a.NewWindow("Make Note")

		// makes a now entry
		FileName := widget.NewEntry()

		// resized the entry
		FileName.Resize(fyne.NewSize(300, FileName.MinSize().Height))

		// makes the submit button
		ButtonSubmit := snippets.MakeButton(snippets.Button{Text: "Make Note", Func: func() {
			// checks if there is no name
			if FileName.Text == "" {
				return
			}

			// checks if the name is more than 15 characters (because it could be too big, and overflow)
			if len(FileName.Text) > 15 {
				// if it is, make a popup saying that it is too big
				popup := widget.NewModalPopUp(
					widget.NewLabel("You cant make it more than 15 characters"),
					MakeNoteWindow.Canvas(),
				)
				popup.Show()
				// wait 2 seconds
				time.Sleep(2 * time.Second)
				// hide the popup, and stop the code for doing anything else
				popup.Hide()
				return
			}

			// checks if the name is already taken
			for _, v := range files.Files {
				if v.Name == FileName.Text {
					// if it is, make a popup saying that it is already taken
					popup := widget.NewModalPopUp(
						widget.NewLabel("That already exists"),
						MakeNoteWindow.Canvas(),
					)
					popup.Show()
					// wait 2 seconds
					time.Sleep(2 * time.Second)
					// hide the popup, and stop the code for doing anything else
					popup.Hide()
					return
				}
			}

			// if it is not, get the name
			name = FileName.Text

			// close the window
			MakeNoteWindow.Close()

			// make a new file
			file := structs.NoteDrawFile{Name: name, LastModified: structs.Date{Month: int(time.Now().Month()), Day: time.Now().Day(), TimeHour: time.Now().Hour(), TimeMin: time.Now().Minute()}}

			// convert the time to 12 hour format
			if file.LastModified.TimeHour > 12 {
				file.LastModified.TimeHour = file.LastModified.TimeHour - 12
			}

			// save the file to a map
			files.Files[name] = file

			// make the card and append it to the container
			ContainerFiles.Add(MakeFileCard(file, files, current, ContainerShowContent, LastContainer))
			ContainerFiles.Objects = snippets.Move(ContainerFiles.Objects, len(ContainerFiles.Objects)-1, 0)
			ContainerFiles.Refresh()
		}})

		// resized the window
		MakeNoteWindow.Resize(fyne.NewSize(float32(300), float32(100)))

		// makes a now container to hold the entry and the submit button
		ContainerForm := container.NewVBox(FileName, ButtonSubmit)

		// set the content of the window to the container
		MakeNoteWindow.SetContent(ContainerForm)

		// show the window
		MakeNoteWindow.Show()
	}})
	return ButtonMakeNote
}
