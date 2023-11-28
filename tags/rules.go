package tags

import (
	"strings"

	"github.com/quasilyte/go-ruleguard/dsl"
	"github.com/quasilyte/go-ruleguard/dsl/types"
)

// https://github.com/quasilyte/go-ruleguard/blob/master/_docs/dsl.md#creating-a-ruleguard-bundle
var Bundle = dsl.Bundle{}

// tags blocks struct tags
func tags(m dsl.Matcher) {
	m.Match("type $s struct{$*_; $f; $*_}").
		Where(m["f"].Text.Matches("(?:[^s]+ +){1,2}`.+`") &&
			m["s"].Filter(tagsFilter)).
		At(m["f"]).
		Report(`Don't use tags`)
}

func tagsFilter(ctx *dsl.VarFilterContext) bool {
	typS := ctx.Type

	asPointer := types.AsPointer(typS)
	if asPointer != nil {
		typS = asPointer.Elem()
	}

	// go-ruleguard не поддерживает циклы, поэтому если будет несколько пакетов, то их нужно прописывать явно && strings.HasPrefix(typS.String(), "ddd/example/app/domain/").
	// Можно посмотреть в сторону codeQL или написание линтера на go 😉
	return strings.HasPrefix(typS.String(), "ddd/example/app/domain/") ||
		strings.HasPrefix(typS.String(), "ddd/testdata/src/tags")
}
