package models

type Book struct {
	ID       string `json:"id" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Author   string `json:"author" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

var Books = []Book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}