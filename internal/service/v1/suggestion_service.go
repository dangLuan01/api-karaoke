package v1service

import (
	"github.com/dangLuan01/karaoke/internal/models"
	"github.com/dangLuan01/karaoke/internal/repository"
	"github.com/dangLuan01/karaoke/internal/utils"
)

type suggestionService struct {
	repo repository.SuggestionRepository
}

func NewSuggestionService(repo repository.SuggestionRepository) SuggestionService {
	return &suggestionService {
		repo: repo,
	}
}

func (ss *suggestionService) SaveSuggestion(search string) error {
	
	found, err := ss.repo.Find(search)
	if err != nil {
		return utils.WrapError(
			string(utils.ErrCodeBadRequest),
			"An error occured find suggestion",
			err,
		)
	}
	if found == nil {
		suggestion := models.SongSuggestion{
			Search: search,
			Count: 1,			
		}
		if err := ss.repo.Store(suggestion); err != nil {
			return utils.WrapError(
				string(utils.ErrCodeBadRequest),
				"An error occured insert suggestion",
				err,
			)
		}
		return nil
	}
	if err := ss.repo.Update(search, *found); err != nil {
		return utils.WrapError(
			string(utils.ErrCodeBadRequest),
			"An error occured update suggestion",
			err,
		)
	}

	return nil
}

func (ss *suggestionService) GetAll() ([]models.SongSuggestion, error) {
	
	ds, err := ss.repo.GetAll()

	if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeBadRequest),
			"An error occured get suggestion",
			err,
		)
	}

	return ds, nil
}