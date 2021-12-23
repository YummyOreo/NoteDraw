package structs

import (
	"fyne.io/fyne/v2/widget"
)

type NoteDrawFile struct {
	Name         string
	LastModified Date
	Content      []NoteType
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
