package repository

import (
	"github.com/dangLuan01/karaoke/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type SqlSongRepository struct {
	songs []models.Song
	db *goqu.Database
}

func NewSqlSongRepository(DB *goqu.Database) SongRepository {
	return &SqlSongRepository{
		songs: make([]models.Song, 0),
		db: DB,
	}
}