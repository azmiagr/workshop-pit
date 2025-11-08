package entity

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	BookID      uuid.UUID `json:"book_id" gorm:"type:varchar(36);primaryKey"`
	Title       string    `json:"title" gorm:"type:varchar(255);not null"`
	Author      string    `json:"author" gorm:"type:varchar(255);not null"`
	Publisher   string    `json:"publisher" gorm:"type:varchar(255);not null"`
	Year        int       `json:"year" gorm:"type:int;not null"`
	ISBN        string    `json:"isbn" gorm:"type:varchar(255);unique;not null"`
	Stock       int       `json:"stock" gorm:"type:int;not null"`
	Description string    `json:"description" gorm:"type:text;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Loans []Loan `json:"loans" gorm:"foreignKey:BookID"`
}
