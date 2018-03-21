package generators

import (
	"fmt"
  "strconv"
  "os"
	"text/template"
  "github.com/madpilot/config_generator/descriptor"
  "github.com/madpilot/config_generator/fields"
)

type CppSourceGenerator struct{}

func defaultValueForBoolean(value bool) string {
  if(value) {
    return "true"
  } else {
    return "false"
  }
}

func defaultValueForUInt16(value uint16) string {
  return strconv.Itoa(int(value))
}

func defaultValueForUInt8(value uint8) string {
  return strconv.Itoa(int(value))
}

func defaultValue(name string, fieldType interface{}) string {
  switch fieldType.(type) {
    case fields.Boolean:
      field := fieldType.(fields.Boolean)
      defaultVal := defaultValueForBoolean(field.Default)
      return fmt.Sprintf("this->%s = %s;", name, defaultVal)
    case fields.UInt16:
      field := fieldType.(fields.UInt16)
      defaultVal := defaultValueForUInt16(field.Default)
      return fmt.Sprintf("this->%s = %s;", name, defaultVal)
    case fields.UInt8:
      field := fieldType.(fields.UInt8)
      defaultVal := defaultValueForUInt8(field.Default)
      return fmt.Sprintf("this->%s = %s;", name, defaultVal)
    case fields.BitMask:
      field := fieldType.(fields.BitMask)
      defaultVal := defaultValueForUInt8(field.Default)
      return fmt.Sprintf("this->%s = %s;", name, defaultVal)
    case fields.String:
      field := fieldType.(fields.String)
      return fmt.Sprintf("this->set_%s(\"%s\");", name, field.Default)
  }
  return fmt.Sprintf("// Unknown Default Type for %s", name)
}

func isString(name string, fieldType interface{}) bool {
  switch fieldType.(type) {
    case fields.String:
      return true;
  }
  return false;
}

func (g CppSourceGenerator) Generate(descriptor *descriptor.Descriptor) {
  funcMap := template.FuncMap{
    "defaultvalue": defaultValue,
    "isstring": isString,
  }
  tmpl, parseError := template.New("cpp_source").Funcs(funcMap).ParseFiles("templates/cpp_source.tmpl")

  if parseError != nil {
		panic(parseError)
	}

	executeErr := tmpl.ExecuteTemplate(os.Stdout, "cpp_source.tmpl", descriptor)
  if executeErr != nil {
		panic(executeErr)
	}
}
