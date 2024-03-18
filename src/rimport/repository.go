package rimport

import "vk-film-library/internal/repository"

type Repository struct {
	Actor repository.Actor
	Movie repository.Movie
	Auth repository.Auth
}

type MockRepository struct {
	Actor *repository.MockActor
	Movie *repository.MockMovie
}
