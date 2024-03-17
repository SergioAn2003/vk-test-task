package movie

import (
	"database/sql"
	"time"
)

type CreateMovieParam struct {
	Title       string    `db:"title"`
	Description string    `db:"description"`
	ReleaseDate time.Time `db:"release_date"`
	Rating      float32   `db:"rating"`
}

func NewCreateMovieParam(title, description string, releaseDate time.Time, rating float32) CreateMovieParam {
	return CreateMovieParam{
		Title:       title,
		Description: description,
		ReleaseDate: releaseDate,
		Rating:      rating,
	}
}

func (c CreateMovieParam) IsValidData() bool {
	return c.Title != "" || c.Description != "" || !c.ReleaseDate.IsZero() || c.Rating > 0
}

type UpdateMovieParam struct {
	ID          int             `db:"movie_id" json:"movie_id"`
	Title       sql.NullString  `db:"title" json:"title"`
	Description sql.NullString  `db:"description" json:"description"`
	ReleaseDate sql.NullTime    `db:"release_date" json:"release_date"`
	Rating      sql.NullFloat64 `db:"rating" json:"rating"`
}
