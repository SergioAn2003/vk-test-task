package rimport

import (
	"grpc-test/config"
	"grpc-test/internal/repository/postgresql"
	"grpc-test/internal/transaction"
	"log"
	"os"
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
			Actors: postgresql.NewActors(),
		},
	}
}
