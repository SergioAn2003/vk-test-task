package actor

import (
	"time"
	"vk-film-library/internal/entity/movie"
)

type Actor struct {
	ID        int       `db:"actor_id" json:"actor_id"`
	Name      string    `db:"name" json:"name"`
	Gender    string    `db:"gender" json:"gender"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
	MovieList []movie.Movie
}

