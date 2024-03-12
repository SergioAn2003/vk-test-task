package postgresql

import (
	"grpc-test/internal/entity/user"
	"grpc-test/internal/repository"
	"grpc-test/internal/transaction"
)

type infoRepository struct{}

func NewInfo() repository.Info {
	return &infoRepository{}
}

func (r *infoRepository) SaveUser(ts transaction.Session, user user.User) error {
	sqlQuery := `
	 insert into users
	 (name, age, is_married)
	 values (:name, :age, :is_married)`

	_, err := SqlxTx(ts).NamedExec(sqlQuery, user)
	return err
}
