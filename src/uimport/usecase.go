package uimport

import (
	"grpc-test/internal/usecase"

	"github.com/sirupsen/logrus"
)

type Usecase struct {
	Actors *usecase.ActorsUsecase
	log    *logrus.Logger
}
