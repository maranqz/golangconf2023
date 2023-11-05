package write

import (
	"strings"

	"github.com/quasilyte/go-ruleguard/dsl"
	"github.com/quasilyte/go-ruleguard/dsl/types"
)

// ruleguard -rules write/rules.go -fix testdata/src/write/main.go

// https://github.com/quasilyte/go-ruleguard/blob/master/_docs/dsl.md#creating-a-ruleguard-bundle
var Bundle = dsl.Bundle{}

const rootPkg = "ddd/testdata/src/write/nested"

// write blocks writing in struct of package.
// It's necessary if we want to prevent changing data in inconsistency way.
func write(m dsl.Matcher) {
	m.Match(
		"$x.$field = $_",
		"$x.$field++",
		"$x.$field += $_",
		"$x.$field--",
		"$x.$field -= $_",
	).
		Where(m["x"].Filter(writeFilter)).
		Do(writeReport)
}

func writeFilter(ctx *dsl.VarFilterContext) bool {
	typX := ctx.Type

	asPointer := types.AsPointer(typX)
	if asPointer != nil {
		typX = asPointer.Elem()
	}

	return strings.HasPrefix(typX.String(), rootPkg)
}

func writeReport(ctx *dsl.DoContext) {
	field := ctx.Var("field")

	ctx.SetReport(ctx.Var("x").Text() + "." + field.Text() + " is readonly")
}
