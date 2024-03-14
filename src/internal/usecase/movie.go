package usecase

import (
	"vk-film-library/bimport"
	"vk-film-library/internal/entity/movie"
	"vk-film-library/internal/transaction"
	"vk-film-library/rimport"

	"github.com/sirupsen/logrus"
)

type MovieUsecase struct {
	log *logrus.Logger
	rimport.RepositoryImports
	*bimport.BridgeImports
}

func NewMovie(log *logrus.Logger, ri rimport.RepositoryImports, bi *bimport.BridgeImports) *MovieUsecase {
	return &MovieUsecase{
		log:               log,
		RepositoryImports: ri,
		BridgeImports:     bi,
	}
}

func CreateMovie(ts transaction.Session, p movie.CreateMovieParam) (err error) {
	// lf := logrus.Fields{"title": p.Title}

	return nil
}
