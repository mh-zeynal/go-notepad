package noteStuff

//an struct that hold title(and also name of the file that must be written),
//date of creating and content of a note
type Note struct {
	Title string	`json:"title"`
	Date string		`json:"date"`
	Content string	`json:"content"`
}

//returns title of a Note
func (n *Note) GetTitle() string {
	return n.Title
}

//returns creating date of a Note
func (n *Note) GetDate() string {
	return n.Date
}

//returns content of a Note
func (n *Note) GetContent() string {
	return n.Content
}
