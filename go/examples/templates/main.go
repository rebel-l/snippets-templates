package main

import (
	"fmt"
	"os"
	"text/template"

	"github.com/rebel-l/snippets-templates/go/examples/templates/tpl"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Init Templates
	tmpl, err := template.New("subject").Parse(tpl.Subject)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	tmpl, err = tmpl.New("salutation").Parse(tpl.Salutation)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	tmpl, err = tmpl.New("salutation-again").Parse(tpl.SalutationAgain)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	tmpl, err = tmpl.New("complementary-close").Parse(tpl.ComplementaryClose)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	tmpl, err = tmpl.New("template-a").Parse(tpl.TemplateA)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	tmpl, err = tmpl.New("template-b").Parse(tpl.TemplateB)
	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
	}

	// Execute specific template
	if err := tmpl.ExecuteTemplate(os.Stdout, "subject", tpl.ParamsA); err != nil {
		log.Fatalf("Error executing template: %s", err)
	}

	fmt.Println() // Just a Spacer

	// Execute specific template
	if err := tmpl.ExecuteTemplate(os.Stdout, "template-a", tpl.ParamsA); err != nil {
		log.Fatalf("Error executing template: %s", err)
	}

	fmt.Println() // Just a Spacer
	fmt.Println("----------------------------")
	fmt.Println()

	// Execute last added template
	if err := tmpl.Execute(os.Stdout, tpl.ParamsB); err != nil {
		log.Fatalf("Error executing template: %s", err)
	}
}
