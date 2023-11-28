package ad

type Attrs map[AttrID]AttrValue

func emptyAttrs() Attrs {
	return make(Attrs)
}

type AttrID int

func NewAttrID(in int) (AttrID, error) {
	// Валидация
	return AttrID(in), nil
}

type AttrValue string

func NewAttrValue(in string) (AttrValue, error) {
	// Валидация
	return AttrValue(in), nil
}
