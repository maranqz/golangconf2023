package nested2

type Nested struct{}

func New() Nested {
	return Nested{}
}

func NewPtr() *Nested {
	return &Nested{}
}
