package uimport

import (
	"grpc-test/bimport"
	"grpc-test/config"
	"grpc-test/internal/transaction"
	"grpc-test/internal/usecase"
	"grpc-test/rimport"
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
	dblog *logrus.Logger,
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
			Info:   usecase.NewInfo(log, dblog, ri, bi),
			Logger: usecase.NewLogger(log, ri),
		},
		BridgeImports: bi,
	}

	return ui
}
