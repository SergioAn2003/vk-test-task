package rimport

import (
	"grpc-test/internal/repository"
)

type Repository struct {
	Info repository.Info
}

type MockRepository struct {
	Info *repository.MockInfo
}
