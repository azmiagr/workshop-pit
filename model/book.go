package model

import (
	"time"

	"github.com/google/uuid"
)

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Publisher   string `json:"publisher" binding:"required"`
	Year        int    `json:"year" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type BookResponse struct {
	BookID      uuid.UUID `json:"book_id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Publisher   string    `json:"publisher"`
	Year        int       `json:"year"`
	ISBN        string    `json:"isbn"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookUpdateRequest struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	Year        int    `json:"year"`
	ISBN        string `json:"isbn"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
}
