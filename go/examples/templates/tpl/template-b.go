package tpl

import (
	"time"

	"github.com/rebel-l/snippets-templates/go/examples/templates/data"
)

// TemplateB ...
var TemplateB = `{{template "salutation-again"}}

On {{.Day}} the {{.Object}} will die.

{{template "complementary-close"}}
`

// ParamsB ...
var ParamsB = struct {
	Day    string
	Object data.Objects
}{
	Day:    time.Now().Weekday().String(),
	Object: data.ListObjects,
}
