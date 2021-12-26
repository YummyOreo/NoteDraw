package exporting

import (
	"NoteDraw/structs"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/harry1453/go-common-file-dialog/cfd"
	"github.com/harry1453/go-common-file-dialog/cfdutil"
)

func Shortcut(files *structs.Files, current *structs.CurrentFile) {
	result, err := cfdutil.ShowSaveFileDialog(cfd.DialogConfig{
		Title: "Save A File",
		Role:  "SaveFileExample",
		FileFilters: []cfd.FileFilter{
			{
				DisplayName: "JSON Files (*.json)",
				Pattern:     "*.json",
			},
			{
				DisplayName: "NoteDraw Files (*.nd)",
				Pattern:     "*.nd",
			},
		},
		SelectedFileFilterIndex: 1,
		FileName:                current.FileName + ".nd",
		DefaultExtension:        "nd",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	if result == "" {
		return
	}

	TempFile := structs.SaveFile{Name: current.FileName, LastModified: files.Files[current.FileName].LastModified}
	for _, v := range files.Files[current.FileName].Content {
		switch v.Type {
		case "title":
			TempFile.Content = append(TempFile.Content, structs.SaveType{Type: v.Type, Data: v.Title.Text.Text})
		case "paragraph":
			TempFile.Content = append(TempFile.Content, structs.SaveType{Type: v.Type, Data: v.Paragraph.Text.Text})
		case "draw":
			lines := []structs.SaveDraw{}
			for _, v := range v.Drawing.Canvas.Line {
				lines = append(lines, structs.SaveDraw{Pos1X: int(v.Position1.X), Pos1Y: int(v.Position1.Y), Pos2X: int(v.Position2.X), Pos2Y: int(v.Position2.Y)})
			}
			TempFile.Content = append(TempFile.Content, structs.SaveType{Type: v.Type, Lines: lines})
		}
	}
	b, err := json.Marshal(TempFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	ioutil.WriteFile(result, b, os.ModePerm)
}
