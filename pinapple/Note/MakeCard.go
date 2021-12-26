package Note

import (
	"NoteDraw/snippets"
	"NoteDraw/structs"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

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
