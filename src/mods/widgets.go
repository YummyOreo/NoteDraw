package mod

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func MakeRect(RectVar Rect) *canvas.Rectangle {

	if RectVar.Color == nil {
		RectVar.Color = color.Black
	}

	rect := canvas.NewRectangle(RectVar.Color)

	if RectVar.StrokeColor != nil {
		rect.StrokeColor = RectVar.StrokeColor
		rect.StrokeWidth = RectVar.StrokeWidth
	}

	return rect
}

type Rect struct {
	Color     color.Color
	FillColor color.Color

	StrokeColor color.Color
	StrokeWidth float32
}

func MakeLine(LineVars Line) *canvas.Line {

	if LineVars.Color == nil {
		LineVars.Color = color.Black
	}

	line := canvas.NewLine(LineVars.Color)

	if LineVars.Position1.X != 0 {
		line.Position1.X = LineVars.Position1.X
		line.Position1.Y = LineVars.Position1.Y

		line.Position2.X = LineVars.Position2.X
		line.Position2.Y = LineVars.Position2.Y
	}

	if LineVars.StrokeColor != nil {
		line.StrokeColor = LineVars.StrokeColor
		line.StrokeWidth = LineVars.StrokeWidth
	}

	return line
}

type Line struct {
	Color     color.Color
	Position1 fyne.Position
	Position2 fyne.Position

	StrokeColor color.Color
	StrokeWidth float32
}

func MakeIcon(IconVars Icon) *widget.Icon {
	if IconVars.Theme == nil {
		IconVars.Theme = theme.AccountIcon()
	}
	icon := widget.NewIcon(IconVars.Theme)
	return icon
}

type Icon struct {
	Theme fyne.Resource
}
