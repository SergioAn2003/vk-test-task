package rimport

import "vk-film-library/internal/repository"

type Repository struct {
	Actor repository.Actor
	Movie repository.Movie
}

type MockRepository struct {
	Actor *repository.MockActor
	Movie *repository.MockMovie
}
