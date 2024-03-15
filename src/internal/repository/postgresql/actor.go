package postgresql

import (
	"vk-film-library/internal/entity/actor"
	"vk-film-library/internal/repository"
	"vk-film-library/internal/transaction"
)

type actorRepository struct{}

func NewActor() repository.Actor {
	return &actorRepository{}
}

func (r *actorRepository) CreateActor(ts transaction.Session, p actor.CreateActorParam) (actorID int, err error) {
	sqlQuery := `
	 insert into actors
	 (name, gender, birth_date)
	 values ($1, $2, $3)
	 returning actor_id`

	err = SqlxTx(ts).QueryRow(sqlQuery, p.Name, p.Gender, p.BirthDate).Scan(&actorID)
	return
}

func (r *actorRepository) Update(ts transaction.Session, p actor.UpdateActorParam) (err error) {
	sqlQuery := `
	update actors set
	name = coalesce(:name, name),
	gender = coalesce(:gender, gender),
	birth_date = coalesce (:birth_date, birth_date)
	where actor_id = :actor_id
	`
	_, err = SqlxTx(ts).NamedExec(sqlQuery, p)
	return
}

func (r *actorRepository) Delete(ts transaction.Session, actorID int) (err error) {
	sqlQuery := `
	delete from actors
	where actor_id = $1`

	_, err = SqlxTx(ts).Exec(sqlQuery, actorID)
	return err
}

func (r *actorRepository) DeleteActorMovie(ts transaction.Session, movieID int) (err error) {
	sqlQuery := `
	delete from actors_movie
	where movie_id = $1`

	_, err = SqlxTx(ts).Exec(sqlQuery, movieID)
	return
}
