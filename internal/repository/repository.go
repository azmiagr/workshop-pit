package repository

import "gorm.io/gorm"

type Repository struct {
	BookRepository IBookRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		BookRepository: NewBookRepository(db),
	}
}
