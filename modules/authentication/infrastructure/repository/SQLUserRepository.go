package repository

import (
	"github.com/franszhafran/sahamify-be/modules/authentication/domain/models"
)

// SQLUserRepository is implementation from it's interface
type SQLUserRepository struct{}

func (s SQLUserRepository) GetByID(id int) bool {
	return true
}

func (s SQLUserRepository) Persist(u models.User) *models.User {
	return &models.User{
		1,
		"a",
		"b",
	}
}

func (s SQLUserRepository) Delete(id int) bool {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return true
}
