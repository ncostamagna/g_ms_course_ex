package course

import (
	"errors"
	"fmt"
)

var ErrInvalidStartDate = errors.New("invalida start date")
var ErrInvalidEndDate = errors.New("invalida end date")
var ErrNameRequired = errors.New("name is required")
var ErrStartRequired = errors.New("start date is required")
var ErrEndRequired = errors.New("end date is required")

// vamos a ir y ver el objeto error de nuevo
type ErrNotFound struct {
	CourseID string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("course '%s' doesn't exist", e.CourseID)
}
