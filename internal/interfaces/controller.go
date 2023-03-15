package interfaces

import (
	"github.com/diyliv/interweb/internal/models"
)

type Controller interface {
	FindInfo(userId int64, category string) ([]models.APIResponse, error)
	GetUserInfo(userId int64) (*models.User, error)
}
