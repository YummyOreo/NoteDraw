package snippets

import (
	"fmt"

	"NoteDraw/structs"
)

func GetPrev(file structs.NoteDrawFile) string {
	file.Content = fmt.Sprintf("%-80v ", file.Content)
	first := file.Content[0:79]
	return first
}
