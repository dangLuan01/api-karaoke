package repository

import (
	"encoding/json"
	"fmt"

	v1dto "github.com/dangLuan01/karaoke/internal/dto/v1"
	"github.com/dangLuan01/karaoke/internal/models"
	"github.com/dangLuan01/karaoke/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

type SqlSongRepository struct {
	songs []models.Song
	db *goqu.Database
	image ImageRepository
}

func NewSqlSongRepository(DB *goqu.Database, image ImageRepository) SongRepository {
	return &SqlSongRepository{
		songs: make([]models.Song, 0),
		db: DB,
		image: image,
	}
}

func (sr * SqlSongRepository) Store(songs []models.Song) error {
	
	_, err := sr.db.Insert(goqu.T("songs")).Rows(songs).Executor().Exec()
	if err != nil {
		return err
	}
	domain := utils.GetEnv("DOMAIN", "")
	var images v1dto.RawImage
	for _, song := range songs {
		url := fmt.Sprintf("%s/%s", domain, song.Id)
		data, err := utils.GetHttpAndDecrypto(url)	

		if err != nil {
			return err
		}

		if err := json.Unmarshal(data, &images); err != nil {
			return err
		}
		
		if err := sr.image.Store(v1dto.MapRawImageToModel(song.Id, song.Uuid, images)); err != nil {
			return err
		}
		
	}
	
	return nil
}

func (sr * SqlSongRepository) FindId(id string) (bool, error) {
	var song models.Song
	ds := sr.db.From(goqu.T("songs").As("s")).
		Where(
			goqu.I("s._id").Eq(id),
		).Limit(1)

	found, err := ds.ScanStruct(&song)
	if err != nil || !found {
		return false, err
	}

	return true, nil
}

func (sr *SqlSongRepository)FindByUuid(uuid string) (*v1dto.SongDTO, error) {
				
	ds := sr.db.From(goqu.T("songs").As("s")).
		Select(
			goqu.I("s.uuid"),
			goqu.I("s.name"),
			goqu.I("s.name_sort"),
			goqu.I("s.singer"),
			goqu.I("s.author"),
			goqu.I("s.tone"),
		).
		LeftJoin(
			goqu.T("song_images").As("is"), goqu.On(
				goqu.I("is.song_uuid").Eq(goqu.I("s.uuid")),
			),
		).
		Where(
			goqu.I("s.uuid").Eq(uuid),
		)
	var (
		song v1dto.SongDTO
		images []v1dto.ImageDTO
	)

	found, err := ds.ScanStruct(&song)

	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("not found")
	}

	ds_images := sr.db.From(goqu.T("song_images").As("si")).
	Join(goqu.T("songs").As("s"), goqu.On(
		goqu.I("si.song_uuid").Eq(goqu.I("s.uuid")),
	)).Where(
		goqu.I("si.song_uuid").Eq(song.Uuid),
	).
	Select(
		goqu.I("si.image"),
	)

	if er := ds_images.ScanStructs(&images); er != nil {
		return nil, er

	}
	song.Images = images

	return &song, nil
}

func (sr *SqlSongRepository) FindByName(name string) ([]models.Song, error) {
	songs := sr.songs[:0]
	ds := sr.db.From(goqu.T("songs").As("s")).
		Select(
			goqu.I("s.uuid"),
			goqu.I("s.name"),
			goqu.I("s.name_sort"),
			goqu.I("s.tone"),
			goqu.I("s.tune"),
			goqu.I("s.singer"),
			goqu.I("s.author"),
			goqu.I("s.thumbnail"),
			goqu.I("s.lyrics"),
			goqu.I("s.type"),
		).
		Where(
			goqu.Or(
				goqu.L("s.name COLLATE utf8mb4_general_ci").Like("%" + name + "%"),
				goqu.I("s.name_sort").Like("%" + name + "%"),
			),
		)
	if err:= ds.ScanStructs(&songs);err != nil {
		return nil, err
	}
	
	return songs, nil
}

func (sr *SqlSongRepository) FindAll() ([]models.Song, error) {
	songs := sr.songs[:0]
	ds := sr.db.From(goqu.T("songs")).Select(
		goqu.C("uuid"),
		goqu.C("name"),
		goqu.C("thumbnail"),
		goqu.C("lyrics"),
		goqu.C("singer"),
		goqu.C("author"),
		goqu.C("tone"),
		goqu.C("tune"),
		goqu.C("type"),
	).
	Order(goqu.I("created_at").Desc()).
	Limit(9)
	if err := ds.ScanStructs(&songs); err != nil {
		return nil, err
	}

	return songs, nil
}