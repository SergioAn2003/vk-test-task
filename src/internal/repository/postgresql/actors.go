package postgresql

import (
	"grpc-test/internal/entity/actor"
	"grpc-test/internal/repository"
	"grpc-test/internal/transaction"
)

type actorsRepository struct{}

func NewActors() repository.Actors {
	return &actorsRepository{}
}

func (r *actorsRepository) CreateActor(ts transaction.Session, actor actor.Actor) error {
	sqlQuery := `
	 insert into actors
	 (name, gender, birth_date)
	 values (:name, :gender, :birth_date)`

	_, err := SqlxTx(ts).NamedExec(sqlQuery, actor)
	return err
}

func (r *actorsRepository) Update(ts transaction.Session, actor actor.Actor) error {
	sqlQuery := `
	update actors set
	name = coalesce(:name, name),
	gender = coalesce(:gender, gender),
	birth_date = coalesce (:birth_date, birth_date)
	where id = :id
	`
	_, err := SqlxTx(ts).NamedExec(sqlQuery, actor)
	return err
}
