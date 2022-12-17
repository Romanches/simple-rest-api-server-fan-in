package errs

import "errors"

// common service errors
var (
	InternalErr             = errors.New("internal error")
	ErrNotValid             = errors.New("invalid parameter")
)
