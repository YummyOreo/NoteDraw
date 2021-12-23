package structs

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type DrawRect struct {
	widget.Card
	Container *fyne.Container
	LineList  *LineList
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
