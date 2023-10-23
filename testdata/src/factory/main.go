package main

import (
	"ddd/testdata/src/factory/creatable"
	"ddd/testdata/src/factory/nested"
	"ddd/testdata/src/factory/nested/nested2"
)

type Struct struct {
}

func main() {
	_ = nested2.Nested{}  // want `Use factory for nested2.Nested`
	_ = &nested2.Nested{} // want `Use factory for nested2.Nested`
	_ = nested2.New()
	_ = nested2.NewPtr()

	_ = nested.Nested{}  // want `Use factory for nested.Nested`
	_ = &nested.Nested{} // want `Use factory for nested.Nested`
	_ = nested.New()
	_ = nested.NewPtr()

	_ = Struct{}
	_ = &Struct{}

	_ = creatable.Struct{}
	_ = &creatable.Struct{}
}
