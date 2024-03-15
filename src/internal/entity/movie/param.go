package movie

import "time"

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
	return c.Title != "" && c.Description != "" && !c.ReleaseDate.IsZero() && c.Rating > 0
}
