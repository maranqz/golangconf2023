//go:build vo_ptr
// +build vo_ptr

package ad

type CountPtr struct {
	available int
	reserved  *int
}

func NewCountPtr(available int) (CountPtr, error) {
	if err := assertAvailable(available); err != nil {
		return CountPtr{}, err
	}

	return CountPtr{
		available: available,
		reserved:  nil,
	}, nil
}

func (c CountPtr) Reserve(r int) (CountPtr, error) {
	if err := assertReserved(r); err != nil {
		return CountPtr{}, err
	}

	available := c.available - r

	if err := assertAvailable(available); err != nil {
		return CountPtr{}, err
	}

	return CountPtr{
		available: available,
		reserved:  &r,
	}, nil
}

func (c CountPtr) IsEqual(another CountPtr) bool {
	return c.available == another.available &&
		(c.reserved == another.reserved ||
			c.reserved != nil && another.reserved != nil &&
				*c.reserved == *another.reserved)
}
