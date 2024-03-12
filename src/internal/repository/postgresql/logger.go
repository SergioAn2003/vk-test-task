package postgresql

import (
	"grpc-test/internal/repository"
	"grpc-test/internal/transaction"
)

type loggerRepository struct{}

func NewLogger() repository.Logger {
	return &loggerRepository{}
}

func (r *loggerRepository) SaveLog(ts transaction.Session, log string) error {
	return nil
}
