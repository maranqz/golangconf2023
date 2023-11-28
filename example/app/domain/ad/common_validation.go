package ad

import (
	"errors"
	"fmt"
)

type Number interface {
	~int | ~int64 //...
}

var (
	ErrNegative = errors.New("negative")
	ErrZero     = errors.New("zero")
)

func assertPositiveBool[In Number](in In) bool {
	return in <= 0
}

func assertPositive[In Number](in In) error {
	if in == 0 {
		return ErrZero
	}

	if in < 0 {
		return ErrNegative
	}

	return nil
}

var (
	ErrLength    = errors.New("length")
	ErrLengthMin = fmt.Errorf("%w: min", ErrLength)
	ErrLengthMax = fmt.Errorf("%w: max", ErrLength)
)

type Measurable interface {
	~string | ~[]any | map[any]any | ~chan<- any
}

func assertLength[In Measurable](in In, min, max int) error {
	ln := len(in)
	if ln < min {
		return ErrLengthMin
	} else if max < ln {
		return ErrLengthMax
	}

	return nil
}

func mapError(err error, mp map[error]error) error {
	if res, ok := mp[err]; ok {
		return res
	}

	return err
}
