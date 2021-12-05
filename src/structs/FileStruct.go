package structs

type NoteDrawFile struct {
	Name         string
	LastModified Date
	Prev         string
}

type Date struct {
	Month    int
	Day      int
	TimeHour int
	TimeMin  int
}
