package v1dto

import (
	"strings"

	"github.com/dangLuan01/karaoke/internal/models"
	"github.com/dangLuan01/karaoke/internal/utils"
	"github.com/google/uuid"
)

type RawSong struct {
	Id 			string `json:"_id"`
	Ten_bai_hat	string `json:"ten_bai_hat"`
	Ca_si   	string `json:"ca_si"`
	Tac_gia   	string `json:"tac_gia"`
	Tone	 	string `json:"tone"`
	Dieu	 	string `json:"dieu"`
	Thumbail 	string `json:"thumbnail"`
	Loi_nhac_goi_y   	string `json:"loi_nhac_goi_y"`
	Kieu_bai_hat	  	string `json:"kieu_bai_hat"`
}

type RawImage struct {
	Id string `json:"_id"`
	Trinh_chieu []string `json:"trinh_chieu"`
	Ten_bai_hat string `json:"ten_bai_hat"`
}

type SongDTO struct {
	Uuid 		string `json:"uuid"`
	Name	  	string `json:"name"`
	Name_sort 	string `json:"name_sort,omitempty"`
	Singer   	string `json:"singer,omitempty"`
	Author   	string `json:"author,omitempty"`
	Tone	 	string `json:"tone,omitempty"`
	Tune	 	string `json:"tune,omitempty"`
	Thumbail 	string `json:"thumbnail,omitempty"`
	Lyrics   	string `json:"lyrics,omitempty"`
	Type	  	string `json:"type,omitempty"`
	Domain_img 		string `json:"domain_img,omitempty"`
	Images 		[]ImageDTO `json:"images,omitempty"`
}

type ImageDTO struct {
	Image string `json:"image"`
}

func MapSongDTO (songs []models.Song) []SongDTO {
	dtos := make([]SongDTO, 0, len(songs))
	for _, song := range songs {
		dto := SongDTO {
			Uuid: song.Uuid,
			Name: song.Name,
			Singer: *song.Singer,
			Author: *song.Author,
			Tone: *song.Tone,
			Tune: *song.Tune,
			Thumbail: *song.Thumbail,
			Lyrics: *song.Lyrics,
			Type: *song.Type,
		}

		dtos = append(dtos, dto)
	}
	return dtos
}

func MapRawSongToModel(songs []RawSong) []models.Song {
	
	modelSong := make([]models.Song, 0, len(songs))
	for _, song := range songs {
		uuid := uuid.New()
		model := models.Song {
			Uuid: uuid.String(),
			Id: song.Id,
			Name: song.Ten_bai_hat,
			Namesort: utils.ToAbbreviation(song.Ten_bai_hat),
			Singer: &song.Ca_si,
			Author: &song.Tac_gia,
			Tone: &song.Tone,
			Tune: &song.Dieu,
			Thumbail: &song.Thumbail,
			Lyrics: &song.Loi_nhac_goi_y,
			Type: &song.Kieu_bai_hat,
		}
		modelSong = append(modelSong, model)
	}

	return  modelSong
}

func MapRawImageToModel(id, uuid string, images RawImage) []models.Image {

	modelImage := make([]models.Image, 0, len(images.Trinh_chieu))

	for _, image := range images.Trinh_chieu {
		parts := strings.SplitN(image, "/", 4)
		path := "/" + parts[3]
		img := models.Image{
			Songuuid: uuid,
			Id: id,
			Image: path,
		}
		modelImage = append(modelImage, img)
	}
	
	return modelImage
}

func MapSongDetailDTO(song *SongDTO) *SongDTO {
	song.Domain_img = utils.GetEnv("DOMAIN_IMG", "")
	return song
}