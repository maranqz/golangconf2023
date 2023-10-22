package bundle_settings

import (
	"github.com/quasilyte/go-ruleguard/dsl"

	"ddd/tags"
)

func init() {
	// Imported rules will have a "qrules" prefix.
	dsl.ImportRules("qrules", tags.Bundle)
}
