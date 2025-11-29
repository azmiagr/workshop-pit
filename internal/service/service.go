package service

import (
	"workshop-pit/internal/repository"
	"workshop-pit/pkg/bcrypt"
	"workshop-pit/pkg/jwt"
)

type Service struct {
	BookService IBookService
	UserService IUserService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) *Service {
	return &Service{
		BookService: NewBookService(repository.BookRepository),
		UserService: NewUserService(repository.UserRepository, bcrypt, jwtAuth),
	}
}
