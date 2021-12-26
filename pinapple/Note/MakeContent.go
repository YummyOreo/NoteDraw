package Note

import (
	"NoteDraw/snippets"
	"NoteDraw/structs"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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

		btn := &structs.DrawRect{LineList: Lines, Container: Box}

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
