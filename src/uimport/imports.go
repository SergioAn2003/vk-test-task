package uimport

import (
	"grpc-test/bimport"
	"grpc-test/config"
	"grpc-test/internal/transaction"
	"grpc-test/internal/usecase"
	"grpc-test/rimport"
	"grpc-test/tools/logger"
	"os"

	"github.com/sirupsen/logrus"
)

type UsecaseImports struct {
	Config         config.Config
	SessionManager transaction.SessionManager
	Usecase        Usecase
	*bimport.BridgeImports
}

func NewUsecaseImports(
	log *logrus.Logger,
	ri rimport.RepositoryImports,
	bi *bimport.BridgeImports,
	sessionManager transaction.SessionManager,
) UsecaseImports {
	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln(err)
	}

	ui := UsecaseImports{
		Config:         config,
		SessionManager: sessionManager,

		Usecase: Usecase{
			Info: usecase.NewInfo(log, ri, bi),
			log:  logger.NewUsecaseLogger(log, "usecase"),
		},
		BridgeImports: bi,
	}

	return ui
}
