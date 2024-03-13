package actor

import (
	"time"
)

type Actor struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Gender    string    `db:"gender" json:"gender"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
}
