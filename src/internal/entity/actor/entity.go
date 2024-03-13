package actor

import (
	"strings"
	"time"
)

type Actor struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Gender    string    `db:"gender" json:"gender"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
}

// IsValidData проверка на валидные значения
func (a Actor) IsValidData() bool {
	return a.Name != "" || (strings.ToLower(a.Gender) == "male" || strings.ToLower(a.Gender) == "female") || !a.BirthDate.Equal(time.Time{})
}
