package usecase

import (
	"vk-film-library/bimport"
	"vk-film-library/internal/entity/global"
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

func (u *MovieUsecase) CreateMovie(ts transaction.Session, p movie.CreateMovieParam) (movieID int, err error) {
	lf := logrus.Fields{"title": p.Title}

	if !p.IsValidData() {
		err = global.ErrParamsIncorect
		return
	}

	movieID, err = u.Repository.Movie.CreateMovie(ts, p)
	if err != nil {
		u.log.WithFields(lf).Errorln("не удалось создать фильм, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	u.log.WithFields(lf).Infof("фильм %s успешно добавлен", p.Title)
	return
}

func (u *MovieUsecase) UpdateMovie(ts transaction.Session, movie movie.Movie) (err error) {
	lf := logrus.Fields{
		"movie_id":         movie.ID,
		"new_title":        movie.Title,
		"new_release_date": movie.ReleaseDate,
		"new_rating":       movie.Rating,
	}

	if err = u.Repository.Movie.UpdateMovie(ts, movie); err != nil {
		u.log.WithFields(lf).Errorln("не удалось обновить данные фильма, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	u.log.WithFields(lf).Infoln("Данные фильма успешно обновлены")
	return
}

func (u *MovieUsecase) DeleteMovie(ts transaction.Session, movieID int) (err error) {
	lf := logrus.Fields{"movie_id": movieID}

	if err = u.Repository.Actor.DeleteActorMovie(ts, movieID); err != nil {
		u.log.WithFields(lf).Errorln("не удалось удалить фильм актера, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	if err = u.Repository.Movie.DeleteMovie(ts, movieID); err != nil {
		u.log.WithFields(lf).Errorln("не удалось удалить фильм, ошибка:", err)
		err = global.ErrInternalError
		return err
	}

	u.log.WithFields(lf).Infoln("фильм успешно удален")
	return
}
