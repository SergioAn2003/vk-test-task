package rimport

import (
	"grpc-test/config"
	"grpc-test/internal/repository/postgresql"
	"grpc-test/internal/repository/redisclient"
	"grpc-test/internal/transaction"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
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

	client := redis.NewClient(&redis.Options{
		Addr:     config.RedisConnectionData().Host,
		Password: config.RedisConnectionData().Password,
		DB:       0,
	})

	return RepositoryImports{
		Config:         config,
		SessionManager: sessionManager,
		Repository: Repository{
			Info:        postgresql.NewInfo(),
			Logger:      postgresql.NewLogger(),
			RedisClient: redisclient.NewRedisClient(client),
		},
	}
}
