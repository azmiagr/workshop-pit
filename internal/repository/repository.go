package repository

import "gorm.io/gorm"

type Repository struct {
	BookRepository IBookRepository
	UserRepository IUserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		BookRepository: NewBookRepository(db),
		UserRepository: NewUserRepository(db),
	}
}
