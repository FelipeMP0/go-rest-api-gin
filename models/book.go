package models

// Book model represents a book with its title and author name, as well as, a unique identifier.
type Book struct {
	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Author string `json:"author"`
}
