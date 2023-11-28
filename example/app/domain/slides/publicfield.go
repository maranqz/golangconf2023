//go:build public_field
// +build public_field

package slides

type Count struct {
	Available int `migration:"available" json:"available"`
	Reserved  int `migration:"reserved" json:"reserved"`

	// или

	available int
	reserved  int
}

func NewCount(a int) (Count, error) {
	return Count{Available: a}, nil
}

func invalidate() {
	q, _ := NewCount(1)
	q.Reserved = -5
}
