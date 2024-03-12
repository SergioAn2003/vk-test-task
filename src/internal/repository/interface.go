package repository

import (
	"grpc-test/internal/entity/user"
	"grpc-test/internal/transaction"
)

type Info interface {
	SaveUser(ts transaction.Session, user user.User) error
}
