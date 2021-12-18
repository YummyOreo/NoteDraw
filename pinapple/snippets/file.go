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

func InsertInt(array []fyne.CanvasObject, value fyne.CanvasObject, index int) []fyne.CanvasObject {
	return append(array[:index], append([]fyne.CanvasObject{value}, array[index:]...)...)
}

func RemoveInt(array []fyne.CanvasObject, index int) []fyne.CanvasObject {
	return append(array[:index], array[index+1:]...)
}

func MoveInt(array []fyne.CanvasObject, srcIndex int, dstIndex int) []fyne.CanvasObject {
	value := array[srcIndex]
	return InsertInt(RemoveInt(array, srcIndex), value, dstIndex)
}
