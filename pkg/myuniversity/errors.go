package myuniversity

import "errors"

var (
	ErrInnoSsoAPIChanged  = errors.New("innopolis SSO API has changed. Needs udpate")
	ErrInvalidCredentials = errors.New("invalid my.university credentials")
)
