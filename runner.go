package config_generator

import (
	"encoding/json"
	"io/ioutil"
  "github.com/madpilot/config_generator/descriptor"
  "github.com/madpilot/config_generator/fields"
  "github.com/madpilot/config_generator/generators"
)

func Generate(file string) {
  data, readError := ioutil.ReadFile(file)

	if readError != nil {
		panic(readError)
	}

	var descriptor descriptor.Descriptor
	jsonErr := json.Unmarshal(data, &descriptor)

	if jsonErr != nil {
		panic(jsonErr)
	}

	for key, field := range descriptor.Fields {
		var newField interface{}

    iface := field.(map[string]interface{})
		fieldType := field.(map[string]interface{})["type"]

		switch fieldType {
		case "bitmask":
			newField = fields.BitMask{
				Default: int(iface["default"].(float64)),
				Length:  int(iface["length"].(float64)),
			}
			break
		case "boolean":
			newField = fields.Boolean{
				Default: iface["default"].(bool),
			}
			break
		case "uint8":
			newField = fields.UInt8{
				Default: uint8(iface["default"].(float64)),
			}
			break
		case "uint16":
			newField = fields.UInt16{
				Default: uint16(iface["default"].(float64)),
			}
			break
		case "string":
			newField = fields.String{
				Default: iface["default"].(string),
			}
			break
		}
		descriptor.Fields[key] = newField
	}

  cppHeader := &generators.CppHeaderGenerator{}
  cppHeader.Generate(&descriptor)
}
