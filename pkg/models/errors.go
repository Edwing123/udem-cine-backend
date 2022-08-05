package models

import "errors"

var (
	ErrAuth             = errors.New("authentication failed")
	ErroNoRows          = errors.New("no rows found")
	ErrServer           = errors.New("server error")
	ErrUserNameTaken    = errors.New("username already taken")
	ErrMovieTitleTaken  = errors.New("movie title already taken")
	ErrRoomTaken        = errors.New("room number already taken")
	ErrScheduleTaken    = errors.New("schedule time already taken")
	ErrFunctionFuckedUp = errors.New("function fucked up")
)
