package repository

import (
	"fmt"
	"time"

	"github.com/dangLuan01/karaoke/internal/models"
	"github.com/doug-martin/goqu/v9"
)

type SqlSuggestionRepository struct {
	suggestion []models.SongSuggestion
	db *goqu.Database
}

func NewSqlSuggestionRepository(DB *goqu.Database) SuggestionRepository {
	return &SqlSuggestionRepository {
		suggestion: make([]models.SongSuggestion, 0),
		db: DB,
	}
}

func (sr *SqlSuggestionRepository) Find(search string) (*models.SongSuggestion, error) {
	var suggestion models.SongSuggestion
	ds := sr.db.From(goqu.T("song_suggestions")).
		Where(
			goqu.C("search").Eq(search),
		)
	found, err := ds.ScanStruct(&suggestion)
	if err != nil {
		return nil, fmt.Errorf("error fetch suggestion:%v", err)
	}

	if !found {
		return nil, nil
	}

	return &suggestion, nil
}

func (sr *SqlSuggestionRepository) Store(suggestion models.SongSuggestion) error {

	insertSuggest := sr.db.Insert("song_suggestions").Rows(suggestion).Executor()
	if _, err := insertSuggest.Exec(); err != nil {
		return fmt.Errorf("error insert suggestion:%v", err)
	}

	return nil
}

func (sr *SqlSuggestionRepository) Update(search string, suggestion models.SongSuggestion) error {
	
	updateSuggestion := sr.db.Update(goqu.T("song_suggestions")).
		Where(
			goqu.C("search").Eq(search),
		).
		Set(goqu.Record{
			"count": suggestion.Count + 1,
			"updated_at": time.Now(),
		}).Executor()
	
	if _, err := updateSuggestion.Exec(); err != nil {
		return fmt.Errorf("error update suggesttion:%v", err)
	}

	return nil
}

func (sr *SqlSuggestionRepository) GetAll() ([]models.SongSuggestion, error) {
	sug := sr.suggestion
	ds := sr.db.From(goqu.T("song_suggestions")).
		Order(
			goqu.C("updated_at").Desc(),
			goqu.C("count").Desc(),
		).Limit(4)
	if err := ds.ScanStructs(&sug); err != nil {
		return nil, fmt.Errorf("error fetch suggesion:%v", err)
	}

	return sug, nil
}
