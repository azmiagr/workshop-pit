package entity

import (
	"time"

	"github.com/google/uuid"
)

type Loan struct {
	LoanID         uuid.UUID `json:"loan_id" gorm:"type:varchar(36);primaryKey"`
	UserID         uuid.UUID `json:"user_id" gorm:"type:varchar(36);not null"`
	BookID         uuid.UUID `json:"book_id" gorm:"type:varchar(36);not null"`
	LoanDate       time.Time `json:"loan_date" gorm:"type:datetime;not null"`
	ReturnDate     time.Time `json:"return_date" gorm:"type:datetime;"`
	MustReturnDate time.Time `json:"must_return_date" gorm:"type:datetime;not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
