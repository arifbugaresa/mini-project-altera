package business

import "errors"

var (
	// ErrInvalidBody Error when data given is not valid on update or insert
	ErrInvalidBody = errors.New("Given body is not valid")
)