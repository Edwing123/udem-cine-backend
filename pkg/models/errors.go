package models

import "errors"

var (
	ErrAuth            = errors.New("authentication failed")
	ErroNoRows         = errors.New("no rows found")
	ErrServer          = errors.New("server error")
	ErrorUserNameTaken = errors.New("username already taken")
)
