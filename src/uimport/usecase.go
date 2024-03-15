package uimport

import (
	"vk-film-library/internal/usecase"

	"github.com/sirupsen/logrus"
)

type Usecase struct {
	Actor *usecase.ActorsUsecase
	Movie *usecase.MovieUsecase
	log   *logrus.Logger
}
