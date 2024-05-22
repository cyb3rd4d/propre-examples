package usecase

import (
	"errors"
	"fmt"
)

var (
	ErrItemNotFound = errors.New("item not found")
)

type ErrInputValidation struct {
	Reason   string
	Previous error
}

func (e ErrInputValidation) Error() string {
	return fmt.Sprintf("input validation error: %s", e.Reason)
}
