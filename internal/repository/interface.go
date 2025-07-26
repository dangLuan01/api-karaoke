package repository

import (
	v1dto "github.com/dangLuan01/karaoke/internal/dto/v1"
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
	FindAll() ([]models.Song, error)
	Store(songs []models.Song) error
	FindId(id string) (bool, error)
	FindByName(name string) ([]models.Song, error)
	FindByUuid(uuid string) (*v1dto.SongDTO, error)
}

type ImageRepository interface {
	Store(images []models.Image) error
}