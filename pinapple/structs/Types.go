package structs

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

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
