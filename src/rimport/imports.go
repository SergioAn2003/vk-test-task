package rimport

import (
	"log"
	"os"
	"vk-film-library/config"
	"vk-film-library/internal/repository/postgresql"
	"vk-film-library/internal/transaction"
)

type RepositoryImports struct {
	Config         config.Config
	SessionManager transaction.SessionManager
	Repository     Repository
}

func NewRepositoryImports(sessionManager transaction.SessionManager) RepositoryImports {
	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln(err)
	}

	return RepositoryImports{
		Config:         config,
		SessionManager: sessionManager,
		Repository: Repository{
			Actor: postgresql.NewActor(),
			Movie: postgresql.NewMovie(),
			Auth:  postgresql.NewAuth(),
		},
	}
}
