package note

import (
	"NoteDraw/snippets"
	"NoteDraw/structs"
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
			ContainerFiles.Objects = snippets.MoveInt(ContainerFiles.Objects, len(ContainerFiles.Objects)-1, 0)
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

func MakeFileCard(file structs.NoteDrawFile, files *structs.Files, current *structs.CurrentFile, ContainerShowContent *fyne.Container, LastContainer *structs.LastContent) *fyne.Container {
	// make a new text, this is so the card is as big as i want it to be
	text := canvas.NewText(fmt.Sprintf("%-80v", ""), color.Gray{Y: 100})

	// makes the card, this shows the name of the file, and the last time it was modified
	card1 := widget.NewCard(files.Files[file.Name].Name, fmt.Sprintf("%02d", files.Files[file.Name].LastModified.Month)+"/"+fmt.Sprintf("%02d", files.Files[file.Name].LastModified.Day)+" "+fmt.Sprintf("%02d", file.LastModified.TimeHour)+":"+fmt.Sprintf("%02d", file.LastModified.TimeMin), text)

	// saves the card to a map, this is so you can change it later
	files.Cards[file.Name] = card1

	// makes the button for opening the file
	btn1 := snippets.MakeButton(snippets.Button{Text: "Open File"})

	// sets the func for the button
	btn1.OnTapped = func() {
		// removes the last content and refreshes the container
		ContainerShowContent.Remove(LastContainer.Content)
		ContainerShowContent.Refresh()

		// adds the new content to the container and refreshes it
		ContainerShowContent.Add(MakeContent(file, files, current, LastContainer))
		ContainerShowContent.Refresh()
	}

	// add them to a container so one is ontop of the other
	VBox1 := container.NewVBox(card1, btn1)
	// returns it so it can be appended to a container
	return container.NewWithoutLayout(VBox1)
}

func MakeContent(file structs.NoteDrawFile, files *structs.Files, current *structs.CurrentFile, LastContainer *structs.LastContent) *container.Scroll {
	// makes and sets a new container to the content of the file
	ContainerContent := snippets.GetContent(files.Files[file.Name], files)

	// sets the current file type to the file types
	current.Types = files.Files[file.Name].Content

	// sets the current file name to the file name
	current.FileName = file.Name

	// makes a new button for making a new paragraph
	btnPh := snippets.MakeButton(snippets.Button{Text: "Paragraph", Func: func() {
		// gets the types of the current file
		NoteTypes := current.Types

		// makes a new multily line text entry
		Ph := widget.NewMultiLineEntry()
		Ph.Wrapping = fyne.TextWrapBreak

		// makes a new type, and sets the text to be the new multiline entry, then appends it to the types
		NoteTypes = append(NoteTypes, structs.NoteType{Type: "paragraph", Paragraph: structs.Paragraph{Text: Ph}})

		// sets the current file types to the new types
		current.Types = NoteTypes

		// adds the paragraph to the container
		ContainerContent.Add(Ph)

		// refreshes the container
		ContainerContent.Refresh()
	}})

	btnTitle := snippets.MakeButton(snippets.Button{Text: "Title", Func: func() {
		// gets the types of the current file
		NoteTypes := current.Types

		// makes a new multily line text entry
		Title := widget.NewEntry()
		Title.TextStyle.Bold = true

		// makes a new type, and sets the text to be the new multiline entry, then appends it to the types
		NoteTypes = append(NoteTypes, structs.NoteType{Type: "title", Title: structs.Title{Text: Title}})

		// sets the current file types to the new types
		current.Types = NoteTypes

		// adds the paragraph to the container
		ContainerContent.Add(Title)

		// refreshes the container
		ContainerContent.Refresh()
	}})

	btnDraw := snippets.MakeButton(snippets.Button{Text: "Draw", Func: func() {

		// gets the types of the current file
		NoteTypes := current.Types

		// makes a new multily line text entry

		Box := container.NewWithoutLayout()

		Lines := new(structs.LineList)

		btn := &snippets.DrawRect{LineList: Lines, Container: Box}

		btn.ExtendBaseWidget(btn)

		rect := snippets.MakeRect(snippets.Rect{Color: color.White})

		rect.SetMinSize(fyne.NewSize(0, 100))

		btn.SetContent(rect)

		// makes a new type, and sets the text to be the new multiline entry, then appends it to the types
		NoteTypes = append(NoteTypes, structs.NoteType{Type: "draw", Drawing: structs.Drawing{Canvas: Lines}})

		// sets the current file types to the new types
		current.Types = NoteTypes

		// adds the paragraph to the container
		ContainerContent.Add(btn)
		ContainerContent.Add(Box)

		// refreshes the container
		ContainerContent.Refresh()

	}})

	// makes it so you can scroll through the content
	ContainerScroll := container.NewScroll(container.NewVBox(container.NewHBox(btnPh, btnTitle, btnDraw), ContainerContent))

	// saves the container to a map, so it can be removed later
	LastContainer.Content = ContainerScroll

	// returns the container
	return ContainerScroll
}
