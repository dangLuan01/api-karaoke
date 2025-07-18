package v1service

import (
	"github.com/dangLuan01/karaoke/internal/models"
)

type UserService interface {
	GetAllUser() ([]models.User, error)
	GetUserByUUID(uuid string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(uuid string, user models.User) (models.User, error)
	DeleteUser(uuid string) error
}

type SongService interface {
	GetAll(data string) (any, error)
}