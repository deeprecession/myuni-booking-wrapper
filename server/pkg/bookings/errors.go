package bookings

import "errors"

var (
	ErrMyUniversityApiChanged = errors.New("my.university API has changed")
	ErrBadCookies             = errors.New("cookies are not valid")
)
