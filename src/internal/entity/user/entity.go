package user

type User struct {
	Login    string `db:"login" json:"login"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
}

func (u User) IsValidRole() bool {
	return u.Role == "admin" || u.Role == "user"
}

type RegisteredUser struct {
	Login    string `db:"login" json:"login"`
	Password []byte `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
}
