package v1service

import (
	v1dto "github.com/dangLuan01/karaoke/internal/dto/v1"
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
	GetAll() ([]models.Song, error)
	SearchSong(search string) ([]models.Song, error)
	GetDetail(uuid string) (*v1dto.SongDTO, error)
}

type SuggestionService interface {
	SaveSuggestion(search string) error
}