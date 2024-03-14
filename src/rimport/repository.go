package rimport

import "vk-film-library/internal/repository"

type Repository struct {
	Actor repository.Actor
}

type MockRepository struct {
	Actor *repository.MockActor
}
