package usecase

import (
	"vk-film-library/bimport"
	"vk-film-library/internal/entity/actor"
	"vk-film-library/internal/entity/global"
	"vk-film-library/internal/transaction"
	"vk-film-library/rimport"

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

func (u *ActorsUsecase) CreateUser(ts transaction.Session, p actor.CreateActorParam) error {
	if !p.IsValidData() {
		return global.ErrParamsIncorect
	}

	if err := u.Repository.Actors.CreateActor(ts, p); err != nil {
		u.log.Errorln("не удалось добавить актёра, ошибка:", err)
		return global.ErrInternalError
	}

	u.log.Infoln("актер успешно добавлен")
	return nil
}
