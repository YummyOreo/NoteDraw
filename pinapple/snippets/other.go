package snippets

import (
	"NoteDraw/structs"
	"fmt"

	"fyne.io/fyne/v2"
	"github.com/kbinani/screenshot"
)

func Insert(array []fyne.CanvasObject, value fyne.CanvasObject, index int) []fyne.CanvasObject {
	return append(array[:index], append([]fyne.CanvasObject{value}, array[index:]...)...)
}

func Remove(array []fyne.CanvasObject, index int) []fyne.CanvasObject {
	return append(array[:index], array[index+1:]...)
}

func Move(array []fyne.CanvasObject, srcIndex int, dstIndex int) []fyne.CanvasObject {
	value := array[srcIndex]
	return Insert(Remove(array, srcIndex), value, dstIndex)
}

func UpdateCard(file structs.NoteDrawFile, files *structs.Files) {
	files.Cards[file.Name].Subtitle = fmt.Sprintf("%02d", files.Files[file.Name].LastModified.Month) + "/" + fmt.Sprintf("%02d", files.Files[file.Name].LastModified.Day) + " " + fmt.Sprintf("%02d", file.LastModified.TimeHour) + ":" + fmt.Sprintf("%02d", file.LastModified.TimeMin)
	files.Cards[file.Name].Refresh()
}

func WindowSize(part float32) fyne.Size {
	if screenshot.NumActiveDisplays() > 0 {
		// #0 is the main monitor
		bounds := screenshot.GetDisplayBounds(0)
		return fyne.NewSize(float32(bounds.Dx())*part, float32(bounds.Dy())*part)
	}
	return fyne.NewSize(800, 800)
}
