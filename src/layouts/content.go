package layouts

import (
	"fyne.io/fyne/v2"
)

type Layout interface {
	Layout([]fyne.CanvasObject, fyne.Size)
	MinSize(objects []fyne.CanvasObject) fyne.Size
}

type TextBox struct {
}

func (d *TextBox) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, o := range objects {
		childSize := o.Size()

		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

func (d *TextBox) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	for _, o := range objects {
		size := o.Size()
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(5, size.Height+10))
	}
}
