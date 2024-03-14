package actor

import (
	"database/sql"
	"strings"
	"time"
)

type CreateActorParam struct {
	Name      string    `db:"name" json:"name"`
	Gender    string    `db:"gender" json:"gender"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
}

func NewCreateActorParam(actorID int, name, gender string, birthDate time.Time) CreateActorParam {
	return CreateActorParam{
		Name:      name,
		Gender:    gender,
		BirthDate: birthDate,
	}
}

func (c CreateActorParam) IsValidData() bool {
	return c.Name != "" && (strings.ToLower(c.Gender) == "male" || strings.ToLower(c.Gender) == "female") && !c.BirthDate.Equal(time.Time{})
}

type UpdateActorParam struct {
	ID        int            `db:"actor_id" json:"actor_id"`
	Name      sql.NullString `db:"name" json:"name"`
	Gender    sql.NullString `db:"gender" json:"gender"`
	BirthDate sql.NullTime   `db:"birth_date" json:"birth_date"`
}

func NewUpdateActorParam(actorID int, name, gender sql.NullString, birthDate sql.NullTime) UpdateActorParam {
	return UpdateActorParam{
		ID:        actorID,
		Name:      name,
		Gender:    gender,
		BirthDate: birthDate,
	}
}
