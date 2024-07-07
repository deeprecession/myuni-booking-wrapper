package myuniversity

import "errors"

var (
	ErrMyUniversityAPIChanged = errors.New("my.univeristy API has changed. Needs udpate")
	ErrInvalidCredentials     = errors.New("invalid my.university credentials")
)
