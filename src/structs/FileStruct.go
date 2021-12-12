package structs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type NoteDrawFile struct {
	Name         string
	LastModified Date
	Prev         string
	Content      string
}

type Date struct {
	Month    int
	Day      int
	TimeHour int
	TimeMin  int
}

type Files struct {
	Files map[string]NoteDrawFile
	Cards map[string]*widget.Card
}

type LastContent struct {
	Content fyne.CanvasObject
}

type CurrentFile struct {
	FileName string
	File     NoteDrawFile
	Card     widget.Card
	Button   widget.Button
}
