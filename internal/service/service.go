package service

import "workshop-pit/internal/repository"

type Service struct {
	BookService IBookService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		BookService: NewBookService(repository.BookRepository),
	}
}
