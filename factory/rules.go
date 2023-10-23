package factory

import (
	"strings"

	"github.com/quasilyte/go-ruleguard/dsl"
)

// ruleguard -rules factory/rules.go -fix testdata/src/factory/main.go

// https://github.com/quasilyte/go-ruleguard/blob/master/_docs/dsl.md#creating-a-ruleguard-bundle
var Bundle = dsl.Bundle{}

const rootPkg = "ddd/testdata/src/factory/nested"

// factory blocks creation of object bypass factory.
func factory(m dsl.Matcher) {
	// Replace on your domain package.
	// To use settings wait answer https://github.com/quasilyte/go-ruleguard/issues/464

	m.Match("$pkg.$x{}").
		Where(!m.File().PkgPath.Matches(rootPkg) &&
			m["pkg"].Object.Is(`PkgName`) &&
			m["x"].Filter(filter)).
		Report(`Use factory for $pkg.$x`)
}

func filter(ctx *dsl.VarFilterContext) bool {
	typX := ctx.Type

	return strings.HasPrefix(typX.String(), rootPkg)
}
