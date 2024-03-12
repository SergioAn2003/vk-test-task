package uimport

import "grpc-test/internal/usecase"

type Usecase struct {
	Info   *usecase.InfoUsecase
	Logger *usecase.LoggerUsecase
}
