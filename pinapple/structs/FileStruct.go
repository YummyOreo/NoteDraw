package structs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type NoteDrawFile struct {
	Name         string
	LastModified Date
	Content      []NoteType
}

type SaveFile struct {
	Name         string
	LastModified Date
	Content      []SaveType
}

type SaveFileJson struct {
	Name         string     `json:"name"`
	LastModified Date       `json:"lastModified"`
	Content      []SaveType `json:"content"`
}

type SaveType struct {
	Type string
	Data string
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
	Drawing   Drawing
}

type Paragraph struct {
	Text *widget.Entry
}

type Title struct {
	Text *widget.Entry
}

type Drawing struct {
	Canvas *LineList
}

type LineList struct {
	Line []*canvas.Line
}
