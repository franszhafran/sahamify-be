package repository

import (
	"github.com/franszhafran/sahamify-be/modules/authentication/domain/models"
)

// SQLUserRepositoryInterface is an interface for user using sql
type SQLUserRepositoryInterface interface {
	GetByID(id int)
	Persist(u models.User) models.User
	Delete(id int) bool
}
