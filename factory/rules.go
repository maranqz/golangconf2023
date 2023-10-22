package factory

import (
	"github.com/quasilyte/go-ruleguard/dsl"
)

// https://github.com/quasilyte/go-ruleguard/blob/master/_docs/dsl.md#creating-a-ruleguard-bundle
var Bundle = dsl.Bundle{}

// factory blocks creation of object bypass factory.
func factory(m dsl.Matcher) {
	// Replace on your domain package.
	// To use settings wait answer https://github.com/quasilyte/go-ruleguard/issues/464
	m.Import("ddd/testdata/src/factory/nested")
	m.Import("ddd/testdata/src/factory/nested/nested2")

	isBlockedPackages := func(v dsl.Var) bool {
		return v.Text == "nested" ||
			v.Text == "nested2"
	}

	m.Match("$pkg.$x{}").
		Where(m["pkg"].Object.Is(`PkgName`) && isBlockedPackages(m["pkg"])).
		Report(`Use factory for $pkg.$x`)
}
