package postgresql

import (
	"vk-film-library/internal/entity/movie"
	"vk-film-library/internal/repository"
	"vk-film-library/internal/transaction"
)

type movieRepository struct{}

func NewMovie() repository.Movie {
	return &movieRepository{}
}

func (r *movieRepository) CreateMovie(ts transaction.Session, p movie.CreateMovieParam) (movieID int, err error) {
	sqlQuery := `
	insert into movies
	(title, description, release_date, rating)
	values ($1, $2, $3, $4)
	returning movie_id`

	err = SqlxTx(ts).QueryRow(sqlQuery, p.Title, p.Description, p.ReleaseDate, p.Rating).Scan(&movieID)
	return
}

func (r *movieRepository) UpdateMovie(ts transaction.Session, movie movie.Movie) (err error) {
	sqlQuery := `
	update movies set
	title = coalesce(:title, title),
	description = coalesce(:description, description),
	release_date = coalesce(:release_date, release_date),
	rating = coalesce(:rating, rating)
	where movie_id = :movie_id`

	_, err = SqlxTx(ts).NamedExec(sqlQuery, movie)
	return
}

func (r *movieRepository) DeleteMovie(ts transaction.Session, movieID int) (err error) {
	sqlQuery := `
	delete from movies
	where movie_id = $1`

	_, err = SqlxTx(ts).Exec(sqlQuery, movieID)
	return
}

