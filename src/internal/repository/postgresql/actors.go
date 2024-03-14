package postgresql

import (
	"vk-film-library/internal/entity/actor"
	"vk-film-library/internal/repository"
	"vk-film-library/internal/transaction"
)

type actorsRepository struct{}

func NewActors() repository.Actors {
	return &actorsRepository{}
}

func (r *actorsRepository) CreateActor(ts transaction.Session, p actor.CreateActorParam) (actorID int, err error) {
	sqlQuery := `
	 insert into actors
	 (name, gender, birth_date)
	 values ($1, $2, $3)
	 returning id`

	err = SqlxTx(ts).QueryRow(sqlQuery, p.Name, p.Gender, p.BirthDate).Scan(&actorID)
	return
}

func (r *actorsRepository) Update(ts transaction.Session, p actor.UpdateActorParam) (err error) {
	sqlQuery := `
	update actors set
	name = coalesce(:name, name),
	gender = coalesce(:gender, gender),
	birth_date = coalesce (:birth_date, birth_date)
	where id = :id
	`
	_, err = SqlxTx(ts).NamedExec(sqlQuery, p)
	return
}

func (r *actorsRepository) Delete(ts transaction.Session, actorID int) (err error) {
	sqlQuery := `
	delete from actors
	where id = $1`

	_, err = SqlxTx(ts).Exec(sqlQuery, actorID)
	return err
}
