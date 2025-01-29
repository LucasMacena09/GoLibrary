package services

import (
	"errors"
	"github.com/LucasMacena09/GoLibrary/models"
)

func GetBookById(id string) (*models.Book, error) {
	for i, b := range models.Books {
		if b.ID == id {
			return &models.Books[i], nil
		}
	}
	return nil, errors.New("book not found")
}
