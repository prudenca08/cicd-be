package bussiness

import "errors"

var (
	ErrEmailorPass    = errors.New("email or password is incorrect")
	ErrDuplicateData  = errors.New("account already exist")
	ErrInternalServer = errors.New("internal server error")
	ErrNotFound       = errors.New("not found")
	ErrUnathorized    = errors.New("unauthorized")

	ErrEmail = errors.New("email doesn't exist")
	ErrPass = errors.New("password doesn't exist")
)
