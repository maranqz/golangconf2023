package ad

import (
	"fmt"
)

var (
	ErrCount          = fmt.Errorf("%w: Quantity", Err)
	ErrCountAvailable = fmt.Errorf("%w: Available should be zero or positive", ErrCount)
	ErrCountReserved  = fmt.Errorf("%w: Reserved should be zero or positive", ErrCount)
)

type Quantity struct {
	// Структурный тег для примера работы линтера
	Available int `json:"available"`
	Reserved  int
}

func DefaultQuantity() Quantity {
	count, _ := NewQuantity(1)

	return count
}

func NewQuantity(available int) (Quantity, error) {
	if err := assertAvailable(available); err != nil {
		return Quantity{}, err
	}

	return Quantity{
		Available: available,
		Reserved:  0,
	}, nil
}

func NewCountWithReserve(available, reserved int) (Quantity, error) {
	c, err := NewQuantity(available)
	if err != nil {
		return Quantity{}, err
	}

	return c.Reserve(reserved)
}

func QuantityFromDB(
	available int,
	reserved int,
	/*...*/
) (Quantity, error) {
	return Quantity{
		Available: available,
		Reserved:  reserved,
	}, nil
}

func (c Quantity) Reserve(r int) (Quantity, error) {
	if err := assertReserved(r); err != nil {
		return Quantity{}, err
	}

	available := c.Available - r

	if err := assertAvailable(available); err != nil {
		return Quantity{}, err
	}

	return Quantity{
		Available: available,
		Reserved:  r,
	}, nil
}

func assertAvailable(available int) error {
	if available < 0 {
		return ErrCountAvailable
	}

	return nil
}

func assertReserved(reserved int) error {
	if reserved < 0 {
		return ErrCountReserved
	}

	return nil
}
