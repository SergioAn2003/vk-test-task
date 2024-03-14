package actor

import (
	"strings"
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

// IsValidData проверка на валидные значения
func (a Actor) IsValidData() bool {
	return a.Name != "" || (strings.ToLower(a.Gender) == "male" || strings.ToLower(a.Gender) == "female") || !a.BirthDate.Equal(time.Time{})
}
