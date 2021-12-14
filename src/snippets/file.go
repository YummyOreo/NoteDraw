package snippets

import (
	"fmt"

	"NoteDraw/structs"
)

func GetPrev(file structs.NoteDrawFile) string {
	if len(file.Content) >= 80 {
		return file.Content[0:80]
	}
	return fmt.Sprintf("%-80v", file.Content)
}
