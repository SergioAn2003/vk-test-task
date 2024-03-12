package usecase

import (
	"grpc-test/bimport"
	"grpc-test/internal/entity/global"
	"grpc-test/internal/entity/user"
	"grpc-test/internal/transaction"
	"grpc-test/rimport"

	"github.com/sirupsen/logrus"
)

type InfoUsecase struct {
	log   *logrus.Logger
	rimport.RepositoryImports
	*bimport.BridgeImports
}

func NewInfo(log *logrus.Logger, ri rimport.RepositoryImports, bi *bimport.BridgeImports) *InfoUsecase {
	return &InfoUsecase{
		log:               log,
		RepositoryImports: ri,
		BridgeImports:     bi,
	}
}

func (u *InfoUsecase) SaveUser(ts transaction.Session, user user.User) error {
	if user.Name == "" {
		return global.ErrInternalError
	}


	if err := u.Repository.Info.SaveUser(ts, user); err != nil {
		u.log.Errorln("не удалось сохранить юзера, ошибка:", err)
		return global.ErrInternalError
	}

	return nil
}
