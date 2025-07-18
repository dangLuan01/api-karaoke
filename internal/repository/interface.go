package repository

import (
	"github.com/dangLuan01/karaoke/internal/models"
)

type UserRepository interface {
	FindAll() ([]models.User, error)
	FindBYUUID(uuid string) (models.User, bool)
	Create(user models.User) error
	Update(uuid string, user models.User) error
	Delete(uuid string) error
	FindByEmail(email string) (models.User, bool)
}

type SongRepository interface {
	
}