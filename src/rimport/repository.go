package rimport

import "vk-film-library/internal/repository"

type Repository struct {
	Actors repository.Actors
}

type MockRepository struct {
	Actors *repository.MockActors
}
