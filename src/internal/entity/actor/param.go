package actor

import (
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
	return c.Name != "" || (strings.ToLower(c.Gender) == "male" || strings.ToLower(c.Gender) == "female") || !c.BirthDate.Equal(time.Time{})
}
