package uimport

import (
	"grpc-test/internal/usecase"

	"github.com/sirupsen/logrus"
)

type Usecase struct {
	Info *usecase.InfoUsecase
	log  *logrus.Logger
}
