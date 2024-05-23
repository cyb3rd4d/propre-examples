package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/cyb3rd4d/poc-propre/internal/item/driver/logger"
)

var (
	ErrItemNotFound = errors.New("item not found")
	ErrInternal     = errors.New("internal error")
)

type ErrInputValidation struct {
	Reason   string
	Previous error
}

func (e ErrInputValidation) Error() string {
	return fmt.Sprintf("input validation error: %s", e.Reason)
}

func logInputValidationError(ctx context.Context, err error) {
	var errInputValidation ErrInputValidation
	if errors.As(err, &errInputValidation) {
		logger.FromContext(ctx).Debug("[use case] input validation error", "error", errInputValidation.Previous)
	}
}
