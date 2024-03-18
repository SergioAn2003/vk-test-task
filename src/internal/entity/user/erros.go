package user

import "errors"

var (
	ErrUnAuthorized = errors.New("неверный логин или пароль")
)