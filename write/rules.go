package write

import (
	"strings"

	"github.com/quasilyte/go-ruleguard/dsl"
	"github.com/quasilyte/go-ruleguard/dsl/types"
)

// https://github.com/quasilyte/go-ruleguard/blob/master/_docs/dsl.md#creating-a-ruleguard-bundle
var Bundle = dsl.Bundle{}

// write blocks writing in struct of package.
// It's necessary if we want to prevent changing data in inconsistency way.
func write(m dsl.Matcher) {
	m.Import("ddd/testdata/src/write/nested")
	m.Import("ddd/testdata/src/write/nested/nested2")

	m.Match(
		"$x.$field = $_",
		"$x.$field++",
		"$x.$field += $_",
		"$x.$field--",
		"$x.$field -= $_",
	).
		Do(report)
}

func report(ctx *dsl.DoContext) {
	typX := ctx.Var("x").Type()

	asPointer := types.AsPointer(typX)
	if asPointer != nil {
		typX = asPointer.Elem()
	}

	if strings.HasPrefix(typX.String(), "ddd/testdata/src/write/nested.") ||
		strings.HasPrefix(typX.String(), "ddd/testdata/src/write/nested/nested2.") {
		field := ctx.Var("field")

		ctx.SetReport(typX.String() + "." + field.Text() + " is readonly")
	}

}
