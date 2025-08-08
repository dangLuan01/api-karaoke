package app

import (
	v1handler "github.com/dangLuan01/karaoke/internal/handler/v1"
	"github.com/dangLuan01/karaoke/internal/repository"
	"github.com/dangLuan01/karaoke/internal/routes"
	v1routes "github.com/dangLuan01/karaoke/internal/routes/v1"
	v1service "github.com/dangLuan01/karaoke/internal/service/v1"
	"github.com/doug-martin/goqu/v9"
)

type SuggestionModule struct {
	routes routes.Route
}

func NewSuggestionModule(DB *goqu.Database) *SuggestionModule {

	suggestionRepo := repository.NewSqlSuggestionRepository(DB)
	suggestionService := v1service.NewSuggestionService(suggestionRepo)
	suggestionHandler := v1handler.NewSuggestionHandler(suggestionService)
	suggestionRoutes := v1routes.NewSuggestionRoutes(suggestionHandler)

	return &SuggestionModule{
		routes: suggestionRoutes,
	}
}
func (s *SuggestionModule) Routes() routes.Route {
	return s.routes
}