package repository

import (
	"grpc-test/internal/entity/actor"
	"grpc-test/internal/transaction"
)

type Actors interface {
	CreateActor(ts transaction.Session, actor actor.Actor) error
}
