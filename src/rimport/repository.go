package rimport

import (
	"grpc-test/internal/repository"
)

type Repository struct {
	Actors repository.Actors
}

type MockRepository struct {
	Actors *repository.MockActors
}
