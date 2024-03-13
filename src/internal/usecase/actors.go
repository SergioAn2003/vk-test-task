package usecase

import (
	"grpc-test/bimport"
	"grpc-test/internal/entity/actor"
	"grpc-test/internal/entity/global"
	"grpc-test/internal/transaction"
	"grpc-test/rimport"

	"github.com/sirupsen/logrus"
)

type ActorsUsecase struct {
	log *logrus.Logger
	rimport.RepositoryImports
	*bimport.BridgeImports
}

func NewActors(log *logrus.Logger, ri rimport.RepositoryImports, bi *bimport.BridgeImports) *ActorsUsecase {
	return &ActorsUsecase{
		log:               log,
		RepositoryImports: ri,
		BridgeImports:     bi,
	}
}

func (u *ActorsUsecase) CreateUser(ts transaction.Session, actor actor.Actor) error {
	if !actor.IsValidData() {
		return global.ErrParamsIncorect
	}

	if err := u.Repository.Actors.CreateActor(ts, actor); err != nil {
		u.log.Errorln("не удалось добавить актёра, ошибка:", err)
		return global.ErrInternalError
	}

	u.log.Infoln("актер успешно добавлен")
	return nil
}
