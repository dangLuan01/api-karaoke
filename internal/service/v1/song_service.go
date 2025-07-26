package v1service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	v1dto "github.com/dangLuan01/karaoke/internal/dto/v1"
	"github.com/dangLuan01/karaoke/internal/models"
	"github.com/dangLuan01/karaoke/internal/repository"
	"github.com/dangLuan01/karaoke/internal/utils"
)

type songService struct {
	repo repository.SongRepository
}

func NewSongService(repo repository.SongRepository) SongService {
	return &songService{
		repo: repo,
	}	
}

func (ss *songService) GetAll() ([]models.Song, error) {
	songs, err := ss.repo.FindAll()
	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeBadRequest),
			"An error occured fetch songs",
			err,
		)
	}

	return songs, nil
}

func (ss *songService) SearchSong(name string) ([]models.Song, error) {

	song, err := ss.repo.FindByName(name)
	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeBadRequest),
			"An error occured find name song",
			err,
		)
	}
	if len(song) > 0 {
		log.Println("from db")
		return song, nil
	}

	var songs []v1dto.RawSong
	domain := utils.GetEnv("DOMAIN", "https://nhacsong.pro/api/client/bai-hat")
	url := fmt.Sprintf("%s?search=%s", domain, url.QueryEscape(name))
	data, err := utils.GetHttpAndDecrypto(url)	
	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeBadRequest),
			"An error occurred get search",
			err,
		)
	}
	if err := json.Unmarshal(data, &songs); err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeBadRequest),
			"An error occurred Unmarshal",
			err,
		)
	}
	var songMaped = []models.Song{}
	for _, song := range songs {
		find, err := ss.repo.FindId(song.Id)
		if err != nil {
			return nil, utils.WrapError(
				string(utils.ErrCodeBadRequest),
				"Find song failed",
				err,
			)
		}
		if !find {
			songMaped = v1dto.MapRawSongToModel(songs)
			if err := ss.repo.Store(songMaped); err != nil {
				return nil, utils.WrapError(
					string(utils.ErrCodeBadRequest),
					"An error occured insert song",
					err,
				)
			}
		}
	}

	return songMaped, nil
}

func (ss *songService) GetDetail(uuid string) (*v1dto.SongDTO, error) {
	song, err := ss.repo.FindByUuid(uuid)
	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeBadRequest),
			"An error occured fetch song",
			err,
		)
	}

	return song, nil
}