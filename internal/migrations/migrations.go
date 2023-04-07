package migrations

import (
	"go-training/app/entities"
	"go-training/internal/database"
)

func init() {
	database.Connection().AutoMigrate(&entities.Book{})
	database.Connection().AutoMigrate(&entities.Author{})
}
