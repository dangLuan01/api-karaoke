package repository

import (
	"github.com/dangLuan01/karaoke/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type SqlImageRepository struct {
	images []models.Image
	db *goqu.Database
}

func NewSqlImageRepository(DB *goqu.Database) ImageRepository {
	return &SqlImageRepository{
		images: make([]models.Image, 0),
		db: DB,
	}
}

func (ir *SqlImageRepository) Store(image []models.Image) error {
	_, err := ir.db.Insert(goqu.T("song_images")).Rows(image).Executor().Exec()
	if err != nil {

		return err
	}

	return nil
}