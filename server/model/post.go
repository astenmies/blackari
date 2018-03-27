package model

// Go’s structs are typed collections of fields.
// They’re useful for grouping data together to form records.

// Post :
type Post struct {
	ID    string
	Slug  string
	Title string
}

type PostResolver struct {
	S *Post
}
