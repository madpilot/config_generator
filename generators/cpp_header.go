package generators

import (
	"os"
	"text/template"
  "github.com/madpilot/config_generator/descriptor"
)

type CppHeaderGenerator struct{}

func (g CppHeaderGenerator) Generate(descriptor *descriptor.Descriptor) {
  tmpl, parseError := template.ParseFiles("templates/cpp_header.tmpl")

  if parseError != nil {
		panic(parseError)
	}

	executeErr := tmpl.ExecuteTemplate(os.Stdout, "cpp_header.tmpl", descriptor.Fields)
	if executeErr != nil {
		panic(executeErr)
	}
}
