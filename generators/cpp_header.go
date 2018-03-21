package generators

import (
	"fmt"
  "os"
	"text/template"
  "github.com/madpilot/config_generator/descriptor"
  "github.com/madpilot/config_generator/fields"
)

type CppHeaderGenerator struct{}

func Type(name string, fieldType interface{}) string {
  switch fieldType.(type) {
    case fields.Boolean:
      return fmt.Sprintf("bool %s;", name)
    case fields.UInt16:
      return fmt.Sprintf("uint16 %s;", name)
    case fields.UInt8:
      return fmt.Sprintf("uint8 %s;", name)
    case fields.BitMask:
      return fmt.Sprintf("uint8 %s;", name)
    case fields.String:
      return fmt.Sprintf("char * %s;", name)
  }
  return fmt.Sprintf("// Unknown Type for %s", name)

}

func Getter(name string, fieldType interface{}) string {
  switch fieldType.(type) {
    case fields.Boolean:
      return fmt.Sprintf("bool get_%s();", name)
    case fields.UInt16:
      return fmt.Sprintf("uint16 get_%s();", name)
    case fields.UInt8:
      return fmt.Sprintf("uint8 get_%s();", name)
    case fields.BitMask:
      return fmt.Sprintf("uint8 get_%s();", name)
    case fields.String:
      return fmt.Sprintf("char * get_%s();", name)
  }
  return fmt.Sprintf("// Unknown Getter Type for %s", name)
}



func Setter(name string, fieldType interface{}) string {
  switch fieldType.(type) {
    case fields.Boolean:
      return fmt.Sprintf("void set_%s(bool val);", name)
    case fields.UInt16:
      return fmt.Sprintf("void set_%s(uint16 val);", name)
    case fields.UInt8:
      return fmt.Sprintf("void set_%s(uint8 val);", name)
    case fields.BitMask:
      return fmt.Sprintf("void set_%s(uint8 val);", name)
    case fields.String:
      return fmt.Sprintf("void set_%s(const char* val);", name)
  }
  return fmt.Sprintf("// Unknown Setter Type for %s", name)
}

func (g CppHeaderGenerator) Generate(descriptor *descriptor.Descriptor) {
  funcMap := template.FuncMap{
    "type": Type,
    "getter": Getter,
    "setter": Setter,
  }
  tmpl, parseError := template.New("cpp_header").Funcs(funcMap).ParseFiles("templates/cpp_header.tmpl")

  if parseError != nil {
		panic(parseError)
	}

	executeErr := tmpl.ExecuteTemplate(os.Stdout, "cpp_header.tmpl", descriptor.Fields)
  if executeErr != nil {
		panic(executeErr)
	}
}
