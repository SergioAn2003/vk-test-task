package repository

import (
	"vk-film-library/internal/entity/actor"
	"vk-film-library/internal/entity/movie"
	"vk-film-library/internal/transaction"
)

type Actor interface {
	CreateActor(ts transaction.Session, p actor.CreateActorParam) (actorID int, err error)
	Update(ts transaction.Session, p actor.UpdateActorParam) (err error)
	Delete(ts transaction.Session, actorID int) (err error)
	DeleteActorMovie(ts transaction.Session, movieID int) (err error)
}

type Movie interface {
	CreateMovie(ts transaction.Session, p movie.CreateMovieParam) (movieID int, err error)
	UpdateMovie(ts transaction.Session, p movie.UpdateMovieParam) (err error)
	DeleteMovie(ts transaction.Session, movieID int) (err error)
	GetMovieList(ts transaction.Session) ([]movie.Movie, error)
	FindMovieListByTitle(ts transaction.Session, title string) ([]movie.Movie, error)
	FindMovieListByActorName(ts transaction.Session, actorName string) ([]movie.Movie, error)
}
