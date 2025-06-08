package auth

import "errors"

var ErrInvalidPassword = errors.New("invalid password")
var ErrUserNotFound = errors.New("user not found")
