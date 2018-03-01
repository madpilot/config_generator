package generators

import (
  "github.com/madpilot/config_generator/descriptor"
)

type Generator interface {
  Generate(descriptor *descriptor.Descriptor)
}
