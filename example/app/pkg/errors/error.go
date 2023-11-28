package errors

import "fmt"

func Must[A1 any](a1 A1, err error) A1 {
	if err != nil {
		panic(err)
	}

	return a1
}

// В проекте использовал такой подход до go 1.20, поэтому у нас появилась функциюя nested. В ней вызывали multierror от uber.
func Nested(parent, child error) error {
	if child == nil {
		return nil
	}

	return fmt.Errorf("%w: %w", parent, child)
}
