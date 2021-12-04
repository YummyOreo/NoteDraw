package structs

type NoteDrawFile struct {
	Name         string
	LastModified Date
	Prev         string
}

type Date struct {
	Month    float32
	Day      float32
	TimeHour float32
	TimeMin  float32
}
