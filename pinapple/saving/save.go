package saving

import (
	"NoteDraw/snippets"
	"NoteDraw/structs"
	"time"
)

func Save(files *structs.Files, current *structs.CurrentFile) {
	// makes a new temp note, that is blank
	TempFile := structs.NoteDrawFile{}

	// sets the types to the same of the current one
	TempFile.Content = current.Types

	// sets the name to the current file name
	TempFile.Name = files.Files[current.FileName].Name

	// sets the last modified to the current time
	TempFile.LastModified = structs.Date{Month: int(time.Now().Month()), Day: time.Now().Day(), TimeHour: time.Now().Hour(), TimeMin: time.Now().Minute()}

	// changes the time to 12 hour format
	if TempFile.LastModified.TimeHour > 12 {
		TempFile.LastModified.TimeHour = TempFile.LastModified.TimeHour - 12
	}

	// updates the map to be the updated one
	files.Files[TempFile.Name] = TempFile

	// updates the card
	snippets.UpdateCard(files.Files[TempFile.Name], files)
}
