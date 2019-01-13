package tpl

import (
	"time"

	"github.com/rebel-l/snippets-templates/go/examples/templates/data"
)

// TemplateB ...
var TemplateB = `{{template "salutation-again"}}

On {{.Day}} the {{.Object}} will die.
{{if .Optional}}Here is something to tell you: {{.Optional}}.
{{end}}
{{template "complementary-close"}}
`

// ParamsB ...
var ParamsB = struct {
	Day      string
	Object   data.Objects
	Optional string
}{
	Day:    time.Now().Weekday().String(),
	Object: data.ListObjects,
}

// ParamsB1 ...
var ParamsB1 = struct {
	Day      string
	Object   data.Objects
	Optional string
}{
	Day:      time.Now().Weekday().String(),
	Object:   data.ListObjects,
	Optional: "This is optional",
}
