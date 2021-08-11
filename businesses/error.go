package businesses

import "errors"

var (
	ErrInternalServer = errors.New("something gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrIDNotFound = errors.New("id not found")

	ErrCategoryNotFound = errors.New("category not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrUsernamePasswordNotFound = errors.New("(Username) or (Password) empty")
)
