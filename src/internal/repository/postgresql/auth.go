package postgresql

import (
	"vk-film-library/internal/entity/user"
	"vk-film-library/internal/repository"
	"vk-film-library/internal/transaction"
	"vk-film-library/tools/gensql"
)

type authRepository struct{}

func NewAuth() repository.Auth {
	return &authRepository{}
}
func (r *authRepository) RegisterUser(ts transaction.Session, user user.RegisteredUser) (err error) {
	sqlQuery := `
	insert into users
	(login, password, role)
	values(:login, :password, :role)`

	_, err = SqlxTx(ts).NamedExec(sqlQuery, user)
	return
}

func (r *authRepository) GetUseInfo(ts transaction.Session, login string) (user.RegisteredUser, error) {
	sqlQuery := `
	select password, role
	from  users
	where login = $1`

	return gensql.Get[user.RegisteredUser](SqlxTx(ts), sqlQuery, login)
}
