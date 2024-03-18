package uimport

import (
	"os"
	"vk-film-library/bimport"
	"vk-film-library/config"
	"vk-film-library/internal/transaction"
	"vk-film-library/internal/usecase"
	"vk-film-library/rimport"
	"vk-film-library/tools/logger"

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
			Actor: usecase.NewActors(log, ri, bi),
			Movie: usecase.NewMovie(log, ri, bi),
			Auth:  usecase.NewAuth(log, ri),
			log:   logger.NewUsecaseLogger(log, "usecase"),
		},
		BridgeImports: bi,
	}

	return ui
}
