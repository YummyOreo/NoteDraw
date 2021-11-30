package mod

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func MakeRect(Color color.Color) fyne.CanvasObject {
	rect := canvas.NewRectangle(Color)
	return rect
}

func MakeLine(Color color.Color, pos1 fyne.Position, pos2 fyne.Position) fyne.CanvasObject {
	fmt.Println("a")

	line := canvas.NewLine(Color)
	line.Position1.X = pos1.X
	line.Position1.Y = pos1.Y

	line.Position2.X = pos2.X
	line.Position2.Y = pos2.Y

	return line
}

func MakeIcon(iconTheme fyne.Resource) fyne.CanvasObject {
	icon := widget.NewIcon(iconTheme)
	return icon
}
