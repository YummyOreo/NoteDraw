package structs

import (
	"fyne.io/fyne/v2"
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

type LastContent struct {
	Content fyne.CanvasObject
}

type CurrentFile struct {
	Types    []NoteType
	FileName string
}

type NoteType struct {
	Type      string
	Paragraph Paragraph
	Title     Title
}

type Paragraph struct {
	Text *widget.Entry
}

type Title struct {
	Text *widget.Entry
}
