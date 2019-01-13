package tpl

import (
	"time"

	"github.com/rebel-l/snippets-templates/go/examples/templates/data"
)

// TemplateA ...
var TemplateA = `{{template "salutation" .}}

This is Variant A!

It shows up a random Joke defined by a parameter: 
"{{.Joke.Rand}}"

{{template "complementary-close"}}
`

// ParamsA ...
var ParamsA = struct {
	Name string
	Date time.Time
	Joke data.Jokes
}{
	Name: "Rebel L",
	Date: time.Now(),
	Joke: data.ListJokes,
}
