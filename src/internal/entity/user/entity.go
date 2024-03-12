package user

type User struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Age       int    `db:"age"`
	IsMarried bool   `db:"is_married"`
}
