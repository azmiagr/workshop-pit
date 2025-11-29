package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID      uuid.UUID `json:"user_id" gorm:"type:varchar(36);primaryKey"`
	RoleID      int       `json:"role_id"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Email       string    `json:"email" gorm:"type:varchar(255);unique;not null"`
	Password    string    `json:"password" gorm:"type:varchar(255);not null"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(255);unique;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Loans []Loan `json:"loans" gorm:"foreignKey:UserID;"`
}
