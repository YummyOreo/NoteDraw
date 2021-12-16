package snippets

import (
	"NoteDraw/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func GetContent(file structs.NoteDrawFile) *fyne.Container {
	content := container.NewVBox()
	for _, v := range file.Content {
		switch v.Type {
		case "title":
			content.AddObject(v.Title.Text)
		case "paragraph":
			content.AddObject(v.Paragraph.Text)
		}
	}
	return content
}
