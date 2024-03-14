package movie

import (
	"database/sql"
)

type Movie struct {
	ID          int             `db:"movie_id" json:"movie_id"`
	Title       sql.NullString  `db:"title"`
	Description sql.NullString  `db:"description"`
	ReleaseDate sql.NullTime    `db:"release_date"`
	Rating      sql.NullFloat64 `db:"rating"`
}
