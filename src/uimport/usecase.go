package uimport

import (
	"vk-film-library/internal/usecase"

	"github.com/sirupsen/logrus"
)

type Usecase struct {
	Actors *usecase.ActorsUsecase
	log    *logrus.Logger
}
