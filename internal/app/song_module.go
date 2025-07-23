package app

import (
	v1handler "github.com/dangLuan01/karaoke/internal/handler/v1"
	"github.com/dangLuan01/karaoke/internal/repository"
	"github.com/dangLuan01/karaoke/internal/routes"
	v1routes "github.com/dangLuan01/karaoke/internal/routes/v1"
	v1service "github.com/dangLuan01/karaoke/internal/service/v1"
	"github.com/doug-martin/goqu/v9"
	
)

type SongModule struct {
	routes routes.Route
}

func NewSongModule(DB *goqu.Database, image repository.ImageRepository) *SongModule {
	songRepo := repository.NewSqlSongRepository(DB, image)
	songService := v1service.NewSongService(songRepo)
	songHandler := v1handler.NewSongHandler(songService)
	songRoutes := v1routes.NewSongRoutes(songHandler)

	return &SongModule{
		routes: songRoutes,
	}
}

func (m *SongModule) Routes() routes.Route {
	return m.routes
}