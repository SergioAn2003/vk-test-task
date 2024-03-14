package repository

import (
	"vk-film-library/internal/entity/actor"
	"vk-film-library/internal/transaction"
)

type Actors interface {
	CreateActor(ts transaction.Session, p actor.CreateActorParam) (actorID int, err error)
	Update(ts transaction.Session, p actor.UpdateActorParam) (err error)
}
