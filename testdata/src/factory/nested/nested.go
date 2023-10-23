package nested

import (
	"ddd/testdata/src/factory/creatable"
	"ddd/testdata/src/factory/nested/nested2"
)

type Nested struct{}

func New() Nested {
	return Nested{}
}

func NewPtr() *Nested {
	return &Nested{}
}

func CallNested2() {
	_ = nested2.Nested{}
	_ = &nested2.Nested{}

	_ = creatable.Struct{}
	_ = &creatable.Struct{}
}
