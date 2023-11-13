package utils

import (
	"errors"
	"fmt"
	"log/slog"
)

type Fn func() (any, error)

func WithAttempts(f Fn, count int, log *slog.Logger) (any, error) {
	for i := 0; i < count; i++ {

		log.Info(fmt.Sprintf("Attempt %d", i))

		a, err := f()
		if err != nil {
			continue
		}

		return a, err
	}

	return nil, errors.New(fmt.Sprintf("after %d attempts not success"))
}
