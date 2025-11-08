package service

import (
	"workshop-pit/entity"
	"workshop-pit/internal/repository"
	"workshop-pit/model"
	"workshop-pit/pkg/database/mariadb"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IBookService interface {
	CreateBook(param model.BookRequest) (*model.BookResponse, error)
	GetAllBooks() ([]*model.BookResponse, error)
	GetBookByID(id string) (*model.BookResponse, error)
	UpdateBook(id string, param model.BookUpdateRequest) (*model.BookResponse, error)
	DeleteBook(id string) error
}

type BookService struct {
	db         *gorm.DB
	repository *repository.Repository
}

func NewBookService(repository *repository.Repository) IBookService {
	return &BookService{
		db:         mariadb.Connection,
		repository: repository,
	}
}

func (s *BookService) CreateBook(param model.BookRequest) (*model.BookResponse, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	bookID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	book := &entity.Book{
		BookID:      bookID,
		Title:       param.Title,
		Author:      param.Author,
		Publisher:   param.Publisher,
		Year:        param.Year,
		ISBN:        param.ISBN,
		Stock:       param.Stock,
		Description: param.Description,
	}

	err = s.repository.BookRepository.CreateBook(tx, book)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	response := &model.BookResponse{
		BookID:      bookID,
		Title:       param.Title,
		Author:      param.Author,
		Publisher:   param.Publisher,
		Year:        param.Year,
		ISBN:        param.ISBN,
		Stock:       param.Stock,
		Description: param.Description,
	}

	return response, nil
}

func (s *BookService) GetAllBooks() ([]*model.BookResponse, error) {
	var response []*model.BookResponse

	books, err := s.repository.BookRepository.GetAllBooks(s.db)
	if err != nil {
		return nil, err
	}

	for _, book := range books {
		response = append(response, &model.BookResponse{
			BookID:    book.BookID,
			Title:     book.Title,
			Author:    book.Author,
			Publisher: book.Publisher,
			Year:      book.Year,
		})
	}

	return response, nil
}

func (s *BookService) GetBookByID(id string) (*model.BookResponse, error) {
	book, err := s.repository.BookRepository.GetBookByID(s.db, id)
	if err != nil {
		return nil, err
	}

	return &model.BookResponse{
		BookID:      book.BookID,
		Title:       book.Title,
		Author:      book.Author,
		Publisher:   book.Publisher,
		Year:        book.Year,
		ISBN:        book.ISBN,
		Stock:       book.Stock,
		Description: book.Description,
	}, nil
}

func (s *BookService) UpdateBook(id string, param model.BookUpdateRequest) (*model.BookResponse, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	book, err := s.repository.BookRepository.GetBookByID(s.db, id)
	if err != nil {
		return nil, err
	}

	book.Title = param.Title
	book.Author = param.Author
	book.Publisher = param.Publisher
	book.Year = param.Year
	book.ISBN = param.ISBN
	book.Stock = param.Stock
	book.Description = param.Description

	err = s.repository.BookRepository.UpdateBook(tx, book)
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	response := &model.BookResponse{
		BookID:      book.BookID,
		Title:       param.Title,
		Author:      param.Author,
		Publisher:   param.Publisher,
		Year:        param.Year,
		ISBN:        param.ISBN,
		Stock:       param.Stock,
		Description: param.Description,
	}

	return response, nil
}

func (s *BookService) DeleteBook(id string) error {
	tx := s.db.Begin()
	defer tx.Rollback()

	err := s.repository.BookRepository.DeleteBook(tx, id)
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}
