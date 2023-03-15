package interfaces

import "github.com/diyliv/interweb/internal/models"

type Repository interface {
	AddUser(user models.User) error
	UserExists(userId int64) bool
	UpdateUser(userId int64) error
	GetUserInfo(userId int64) (*models.User, error)
}
