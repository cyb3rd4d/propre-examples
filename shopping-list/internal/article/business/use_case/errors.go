package usecase

import (
	"context"
	"errors"
	"fmt"

	"shopping-list/internal/article/driver/logger"
)

var (
	ErrArticleNotFound = errors.New("article not found")
	ErrInternal        = errors.New("internal error")
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
