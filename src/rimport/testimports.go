package rimport

import (
	"grpc-test/config"
	"grpc-test/internal/repository"
	"grpc-test/internal/transaction"
	"log"
	"os"

	"go.uber.org/mock/gomock"
)

type TestRepositoryImports struct {
	Config         config.Config
	SessionManager *transaction.MockSessionManager
	MockRepository MockRepository
	ctrl           *gomock.Controller
}

func NewTestRepositoryImports(
	ctrl *gomock.Controller,
) TestRepositoryImports {
	config, err := config.NewConfig(os.Getenv("CONF_PATH"))
	if err != nil {
		log.Fatalln(err)
	}

	return TestRepositoryImports{
		ctrl:           ctrl,
		Config:         config,
		SessionManager: transaction.NewMockSessionManager(ctrl),
		MockRepository: MockRepository{
			Info:        repository.NewMockInfo(ctrl),
			Logger:      repository.NewMockLogger(ctrl),
			RedisClient: repository.NewMockRedisClient(ctrl),
		},
	}
}

func (t *TestRepositoryImports) MockSession() *transaction.MockSession {
	ts := transaction.NewMockSession(t.ctrl)

	ts.EXPECT().Start().Return(nil).AnyTimes()
	ts.EXPECT().Rollback().Return(nil).AnyTimes()

	return ts
}

func (t *TestRepositoryImports) MockSessionWithCommit() *transaction.MockSession {
	ts := t.MockSession()

	ts.EXPECT().Commit().Return(nil).AnyTimes()

	return ts
}

func (t *TestRepositoryImports) RepositoryImports() RepositoryImports {
	return RepositoryImports{
		SessionManager: t.SessionManager,
		Config:         t.Config,
		Repository: Repository{
			Info:        t.MockRepository.Info,
			Logger:      t.MockRepository.Logger,
			RedisClient: t.MockRepository.RedisClient,
		},
	}
}
