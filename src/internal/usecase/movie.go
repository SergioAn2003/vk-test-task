package usecase

import (
	"sort"
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

func (u *MovieUsecase) UpdateMovie(ts transaction.Session, p movie.UpdateMovieParam) (err error) {
	lf := logrus.Fields{
		"movie_id":         p.ID,
		"new_title":        p.Title,
		"new_release_date": p.ReleaseDate,
		"new_rating":       p.Rating,
	}

	if err = u.Repository.Movie.UpdateMovie(ts, p); err != nil {
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

func (u *MovieUsecase) GetMovieList(ts transaction.Session, sortBy string) ([]movie.Movie, error) {
	lf:=logrus.Fields{"sort_by":sortBy}

	movieList, err := u.Repository.Movie.GetMovieList(ts)
	switch err {
	case nil:
	case global.ErrNoData:
		u.log.WithFields(lf).Debugln("нет фильмов в базе")
		return nil, err
	default:
		u.log.WithFields(lf).Errorln("не удалось получить фильмы, ошибка:", err)
		return nil, global.ErrInternalError
	}

	switch sortBy {
	case "title":
		u.SortMovieListByTitle(movieList)
	case "release_date":
		u.SortMovieListByReleaseDate(movieList)
	default:
		u.SortMovieListByRating(movieList)
	}

	return movieList, nil
}

func (u *MovieUsecase) FindMovieListByTitleAndActorName(ts transaction.Session, title, actorName string) (movieList []movie.Movie, err error) {
	lf := logrus.Fields{
		"movie_titile_fragment": title,
		"actor_name_fragment":   actorName,
	}

	if title != "" && actorName == "" {
		movieList, err = u.Repository.Movie.FindMovieListByTitle(ts, title)
		switch err {
		case nil:
		case global.ErrNoData:
			u.log.WithFields(lf).Debugln("нет фильмов с таким названием")
			return
		default:
			u.log.WithFields(lf).Errorln("не удалось найти фильмы, ошибка:", err)
			err = global.ErrInternalError
			return
		}
	} else if title == "" && actorName != "" {
		movieList, err = u.Repository.Movie.FindMovieListByActorName(ts, actorName)
		switch err {
		case nil:
		case global.ErrNoData:
			u.log.WithFields(lf).Debugln("нет фильмов по этому имени актера")
			return
		default:
			u.log.WithFields(lf).Errorln("не удалось найти фильмы, ошибка:", err)
			err = global.ErrInternalError
			return
		}
	} else {
		err = global.ErrParamsIncorect
		return
	}

	return
}

func (u *MovieUsecase) SortMovieListByRating(movieList []movie.Movie) {
	sort.Slice(movieList, func(i, j int) bool {
		return movieList[i].Rating > movieList[j].Rating
	})
}

func (u *MovieUsecase) SortMovieListByTitle(movieList []movie.Movie) {
	sort.Slice(movieList, func(i, j int) bool {
		return movieList[i].Title < movieList[j].Title
	})
}

func (u *MovieUsecase) SortMovieListByReleaseDate(movieList []movie.Movie) {
	sort.Slice(movieList, func(i, j int) bool {
		return movieList[i].ReleaseDate.Before(movieList[j].ReleaseDate)
	})
}
