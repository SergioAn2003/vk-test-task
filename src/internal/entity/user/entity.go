package user

import (
	"database/sql"
)

type User struct {
	ID        int           `db:"id"`
	Name      string        `db:"name"`
	Age       sql.NullInt64 `db:"age"`
	IsMarried bool          `db:"is_married"`
}
