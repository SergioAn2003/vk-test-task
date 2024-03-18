package movie

import "time"

type Movie struct {
	ID          int       `db:"movie_id" json:"movie_id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	ReleaseDate time.Time `db:"release_date" json:"release_date"`
	Rating      float64   `db:"rating" json:"rating"`
}
