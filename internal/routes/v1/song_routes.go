package v1routes

import (
	v1handler "github.com/dangLuan01/karaoke/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type SongRoutes struct {
	handler *v1handler.SongHandler
}

func NewSongRoutes(handler *v1handler.SongHandler) *SongRoutes {
	return &SongRoutes{
		handler: handler,
	}
}

func (rs *SongRoutes)Register(r *gin.RouterGroup)  {
	song := r.Group("/song")
	{
		// song.GET("", rs.handler.GetAll)
		song.GET("/search", rs.handler.SearchSong)
		song.GET("/:uuid", rs.handler.GetDetail)
	}


}