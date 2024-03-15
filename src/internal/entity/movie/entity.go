package movie

import (
	"database/sql"
)

type Movie struct {
	ID          int             `db:"movie_id" json:"movie_id"`
	Title       sql.NullString  `db:"title" json:"title"`
	Description sql.NullString  `db:"description" json:"description"`
	ReleaseDate sql.NullTime    `db:"release_date" json:"release_date"`
	Rating      sql.NullFloat64 `db:"rating" json:"rating"`
}
