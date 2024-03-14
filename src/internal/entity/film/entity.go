package film

type Film struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
