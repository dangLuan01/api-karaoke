package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/karaoke/internal/dto/v1"
	v1service "github.com/dangLuan01/karaoke/internal/service/v1"
	"github.com/dangLuan01/karaoke/internal/utils"
	"github.com/dangLuan01/karaoke/internal/validation"
	"github.com/gin-gonic/gin"
)

type SongHandler struct {
	service v1service.SongService
}
type Search struct {
	Search string `form:"search"`
}
type GetSongUuidParam struct {
	Uuid string `uri:"uuid" binding:"uuid"`
}

func NewSongHandler(service v1service.SongService) *SongHandler {
	return &SongHandler{
		service: service,
	}
}

func (sh *SongHandler) GetAll(ctx *gin.Context) {
	songs, err := sh.service.GetAll()
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}
	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapSongDTO(songs))
}

func (sh *SongHandler) SearchSong(ctx *gin.Context)  {
	var param Search

	err := ctx.ShouldBindQuery(&param)

	if err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))

		return 
	}
	songs, err := sh.service.SearchSong(param.Search)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapSongDTO(songs))
}

func (sh *SongHandler) GetDetail(ctx *gin.Context) {
	var params GetSongUuidParam

	if err := ctx.ShouldBindUri(&params); err != nil {
		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}

	song, err := sh.service.GetDetail(params.Uuid)
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapSongDetailDTO(song))
	
}