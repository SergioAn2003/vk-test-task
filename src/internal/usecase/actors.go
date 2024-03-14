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

func (u *ActorsUsecase) CreateActor(ts transaction.Session, p actor.CreateActorParam) (actorID int, err error) {
	if !p.IsValidData() {
		err = global.ErrParamsIncorect
		return
	}

	actorID, err = u.Repository.Actors.CreateActor(ts, p)
	if err != nil {
		u.log.Errorln("не удалось добавить актера, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	u.log.Infoln("актер успешно добавлен")
	return
}

func (u *ActorsUsecase) UpdateActor(ts transaction.Session, p actor.UpdateActorParam) (err error) {
	if err = u.Repository.Actors.Update(ts, p); err != nil {
		u.log.Errorln("не удалось обновить данные актера, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	u.log.Infoln("данные актера успешно обновлены")
	return
}
