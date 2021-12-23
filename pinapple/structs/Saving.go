package structs

type SaveFile struct {
	Name         string
	LastModified Date
	Content      []SaveType
}

type SaveFileJson struct {
	Name         string     `json:"name"`
	LastModified Date       `json:"lastModified"`
	Content      []SaveType `json:"content"`
}

type SaveType struct {
	Type  string
	Data  string
	Lines []SaveDraw
}

type SaveDraw struct {
	Pos1X int
	Pos1Y int
	Pos2X int
	Pos2Y int
}
