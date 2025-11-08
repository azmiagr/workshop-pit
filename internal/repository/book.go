package repository

import (
	"workshop-pit/entity"

	"gorm.io/gorm"
)

type IBookRepository interface {
	CreateBook(tx *gorm.DB, book *entity.Book) error
	GetAllBooks(tx *gorm.DB) ([]*entity.Book, error)
	GetBookByID(tx *gorm.DB, id string) (*entity.Book, error)
	UpdateBook(tx *gorm.DB, book *entity.Book) error
	DeleteBook(tx *gorm.DB, id string) error
}

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) IBookRepository {
	return &BookRepository{db: db}
}

func (r *BookRepository) CreateBook(tx *gorm.DB, book *entity.Book) error {
	err := tx.Create(book).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) GetAllBooks(tx *gorm.DB) ([]*entity.Book, error) {
	var books []*entity.Book
	err := tx.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (r *BookRepository) GetBookByID(tx *gorm.DB, id string) (*entity.Book, error) {
	var book *entity.Book
	err := tx.Where("id = ?", id).First(&book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookRepository) UpdateBook(tx *gorm.DB, book *entity.Book) error {
	err := tx.Updates(book).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) DeleteBook(tx *gorm.DB, id string) error {
	err := tx.Delete(&entity.Book{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
