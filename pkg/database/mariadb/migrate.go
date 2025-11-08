package mariadb

import (
	"workshop-pit/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&entity.User{},
		&entity.Book{},
		&entity.Loan{},
	)

	if err != nil {
		return err
	}

	return nil
}
