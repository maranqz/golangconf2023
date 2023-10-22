package tags

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

// https://github.com/quasilyte/go-ruleguard/blob/master/_docs/dsl.md#creating-a-ruleguard-bundle
var Bundle = dsl.Bundle{}

// tags blocks struct tags
func tags(m dsl.Matcher) {
	m.Match("type $_ struct{$*_; $x; $*_}").
		Where(m["x"].Text.Matches("(?:[^s]+ +){1,2}`.+`")).
		At(m["x"]).
		Report(`Don't use tags`)
}
