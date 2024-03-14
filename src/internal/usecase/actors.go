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
	lf := logrus.Fields{
		"name":       p.Name,
		"gender":     p.Gender,
		"birth_date": p.BirthDate,
	}

	if !p.IsValidData() {
		err = global.ErrParamsIncorect
		return
	}

	actorID, err = u.Repository.Actors.CreateActor(ts, p)
	if err != nil {
		u.log.WithFields(lf).Errorln("не удалось добавить актера, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	lf["actor_id"] = actorID

	u.log.WithFields(lf).Infoln("актер успешно добавлен")
	return
}

func (u *ActorsUsecase) UpdateActor(ts transaction.Session, p actor.UpdateActorParam) (err error) {
	lf := logrus.Fields{
		"actor_id":       p.ID,
		"new_name":       p.Name,
		"new_gender":     p.Gender,
		"new_birth_date": p.BirthDate,
	}

	if err = u.Repository.Actors.Update(ts, p); err != nil {
		u.log.WithFields(lf).Errorln("не удалось обновить данные актера, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	u.log.WithFields(lf).Infoln("данные актера успешно обновлены")
	return
}

func (u *ActorsUsecase) DeleteActor(ts transaction.Session, actorID int) (err error) {
	lf := logrus.Fields{"actor_id": actorID}

	if actorID <= 0 {
		err = global.ErrParamsIncorect
		return
	}

	if err = u.Repository.Actors.Delete(ts, actorID); err != nil {
		u.log.WithFields(lf).Errorln("не удалось удалить актера, ошибка:", err)
		err = global.ErrInternalError
		return
	}

	u.log.WithFields(lf).Infoln("актер успешно удален")
	return
}
