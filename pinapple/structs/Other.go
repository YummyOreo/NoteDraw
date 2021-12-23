package structs

import "fyne.io/fyne/v2"

type LastContent struct {
	Content fyne.CanvasObject
}

type CurrentFile struct {
	Types    []NoteType
	FileName string
}
