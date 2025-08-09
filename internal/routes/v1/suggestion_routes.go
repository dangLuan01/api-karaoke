package v1routes

import (
	v1handler "github.com/dangLuan01/karaoke/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type SuggestionRoutes struct {
	handler *v1handler.SuggestionHandler
}

func NewSuggestionRoutes(handler *v1handler.SuggestionHandler) *SuggestionRoutes {
	return &SuggestionRoutes{
		handler: handler,
	}
}

func (rs *SuggestionRoutes) Register(r *gin.RouterGroup) {
	song := r.Group("/suggestion")
	{
		song.POST("", rs.handler.SaveSuggestionBySearch)
		song.GET("/list", rs.handler.GetAllSuggestion)
	}

}