package v1handler

import (
	"net/http"

	v1dto "github.com/dangLuan01/karaoke/internal/dto/v1"
	v1service "github.com/dangLuan01/karaoke/internal/service/v1"
	"github.com/dangLuan01/karaoke/internal/utils"
	"github.com/dangLuan01/karaoke/internal/validation"
	"github.com/gin-gonic/gin"
)

type SuggestionHandler struct {
	service v1service.SuggestionService
}

func NewSuggestionHandler(service v1service.SuggestionService) *SuggestionHandler {
	return &SuggestionHandler {
		service: service,
	}
}

func (sh *SuggestionHandler) SaveSuggestionBySearch(ctx *gin.Context)  {
	var input v1dto.CreateSuggestionInput
	if err := ctx.ShouldBindJSON(&input); err != nil {

		utils.ResponseValidator(ctx, validation.HandlerValidationErrors(err))
		return
	}
	if err := sh.service.SaveSuggestion(input.Search); err != nil {

		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSatus(ctx, http.StatusNoContent)
}

func (sh *SuggestionHandler)GetAllSuggestion(ctx *gin.Context) {
	
	sug, err := sh.service.GetAll()
	if err != nil {
		utils.ResponseError(ctx, err)
		return
	}

	utils.ResponseSuccess(ctx, http.StatusOK, v1dto.MapSuggestionDTO(sug))
	
}