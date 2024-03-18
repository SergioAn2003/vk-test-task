package postgresql

import (
	"vk-film-library/internal/entity/movie"
	"vk-film-library/internal/repository"
	"vk-film-library/internal/transaction"
	"vk-film-library/tools/gensql"
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

func (r *movieRepository) UpdateMovie(ts transaction.Session, p movie.UpdateMovieParam) (err error) {
	sqlQuery := `
	update movies set
	title = coalesce(:title, title),
	description = coalesce(:description, description),
	release_date = coalesce(:release_date, release_date),
	rating = coalesce(:rating, rating)
	where movie_id = :movie_id`

	_, err = SqlxTx(ts).NamedExec(sqlQuery, p)
	return
}

func (r *movieRepository) DeleteMovie(ts transaction.Session, movieID int) (err error) {
	sqlQuery := `
	delete from movies
	where movie_id = $1`

	_, err = SqlxTx(ts).Exec(sqlQuery, movieID)
	return
}

func (r *movieRepository) GetMovieList(ts transaction.Session) ([]movie.Movie, error) {
	sqlQuery := `
	select movie_id, title, description, release_date, rating
	from movies`

	return gensql.Select[movie.Movie](SqlxTx(ts), sqlQuery)
}

func (r *movieRepository) FindMovieListByTitle(ts transaction.Session, title string) ([]movie.Movie, error) {
	sqlQuery := `
	select movie_id, title, description, release_date, rating
	from movies
	where lower(title) like '%' || lower($1) || '%'`

	return gensql.Select[movie.Movie](SqlxTx(ts), sqlQuery, title)
}

func (r *movieRepository) FindMovieListByActorName(ts transaction.Session, actorName string) ([]movie.Movie, error) {
	sqlQuery := `
	select m.movie_id, m.title, m.description, m.release_date, m.rating
	from movies m
	join actors_movie am on (am.movie_id = m.movie_id)
	join actors a on (am.actor_id = a.actor_id)
	where lower(a.name) like '%' || lower($1) || '%'`

	return gensql.Select[movie.Movie](SqlxTx(ts), sqlQuery, actorName)
}
