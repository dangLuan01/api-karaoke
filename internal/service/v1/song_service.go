package v1service

import (
	"fmt"

	"github.com/Luzifer/go-openssl"
	"github.com/dangLuan01/karaoke/internal/repository"
	"github.com/dangLuan01/karaoke/internal/utils"
)

type songService struct {
	repo repository.SongRepository
}
var Song map[string]interface{} 
func NewSongService(repo repository.SongRepository) SongService {
	return &songService{
		repo: repo,
	}	
}

func (ss *songService) GetAll(data, secret string) (any, error) {
	
    o := openssl.New()

    dec, err := o.DecryptBytes(secret, []byte(data))
    if err != nil {
		return nil, utils.WrapError(
			string(utils.ErrCodeBadRequest),
			"An error occurred DecryptBytes",
			err,
		)
    }
	
	// if err := json.Unmarshal(dec, &Song); err != nil {
	// 	return nil, utils.WrapError(
	// 		string(utils.ErrCodeBadRequest),
	// 		"An error occurred Unmarshal",
	// 		err,
	// 	)
	// }
    
	return fmt.Sprintf("Decrypted text: %s\n", string(dec)), nil
}