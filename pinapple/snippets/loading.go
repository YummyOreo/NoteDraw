package snippets

import (
	"NoteDraw/structs"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

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

			btn := &structs.DrawRect{Container: Box, LineList: v.Drawing.Canvas}

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
