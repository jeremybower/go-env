package env

import (
	"errors"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func NotEmpty(v string) error {
	if strings.TrimSpace(v) == "" {
		return errors.New("value must not be empty")
	}

	return nil
}

func NonNegative[T Number](v T) error {
	if v < T(0) {
		return errors.New("value must be non-negative")
	}

	return nil
}

func Positive[T Number](v T) error {
	if v <= T(0) {
		return errors.New("value must be positive")
	}

	return nil
}
