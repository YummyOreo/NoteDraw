package snippets

import (
	"NoteDraw/structs"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type DrawRect struct {
	widget.Card
	Container *fyne.Container
	LineList  *structs.LineList
	fyne.Draggable
	First fyne.Position
	End   fyne.Position
}

func (b *DrawRect) Dragged(d *fyne.DragEvent) {

	if (b.Card.Position().X+b.Card.Size().Width)-15 < d.Position.X {
		return
	}
	if (b.Card.Size().Height)-22 < d.Position.Y {
		return
	}
	if (b.Card.Position().X)+5 > d.Position.X {
		return
	}
	if (b.Card.Size().Height - (b.Card.Size().Height)) > d.Position.Y {
		return
	}

	if b.First.X == 0 {
		b.First = d.Position
		return
	}

	b.End = d.Position
}

func (e *DrawRect) DragEnd() {
	line1 := canvas.NewLine(color.Black)
	line1.StrokeWidth = 2

	line1.Position1 = e.First.Subtract(fyne.NewPos(0, 110))
	line1.Position2 = e.End.Subtract(fyne.NewPos(0, 110))

	e.First = fyne.NewPos(0, 0)
	e.End = fyne.NewPos(0, 0)

	e.LineList.Line = append(e.LineList.Line, line1)

	e.Container.Add(line1)
	e.Container.Refresh()
}

func GetContent(file structs.NoteDrawFile, files *structs.Files) *fyne.Container {
	content := container.NewVBox()
	for _, v := range file.Content {
		switch v.Type {
		case "title":
			content.Add(v.Title.Text)
		case "paragraph":
			content.Add(v.Paragraph.Text)
		case "draw":
			Box := container.NewWithoutLayout()

			btn := &DrawRect{Container: Box, LineList: v.Drawing.Canvas}

			btn.ExtendBaseWidget(btn)

			rect := MakeRect(Rect{Color: color.White})

			rect.SetMinSize(fyne.NewSize(0, 100))

			btn.SetContent(rect)

			for _, i := range v.Drawing.Canvas.Line {
				line1 := canvas.NewLine(color.Black)
				line1.StrokeWidth = 2

				line1.Position1 = i.Position1
				line1.Position2 = i.Position2

				Box.Add(line1)
				Box.Refresh()
			}

			// adds the paragraph to the container
			content.Add(btn)
			content.Add(Box)
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
