package builder

import (
	"fmt"
	"strings"
)

type ResourceType struct {
	Name       string                 `yaml:"name,omitempty" validate:"required"` // TODO: add validation later
	Type       string                 `yaml:"type,omitempty" validate:"required"`
	Source     map[string]interface{} `yaml:"source,omitempty"`
	Privileged bool                   `yaml:"privileged,omitempty"`
	Params     map[string]interface{} `yaml:"params,omitempty"`
	CheckEvery string                 `yaml:"check_every,omitempty"`
	Tags       []string               `yaml:"tags,omitempty"`
}

type ResourceTypes []ResourceType

func (r ResourceType) Generate() string {
	var parts = []string{
		"ResourceType{", // placeholder
		fmt.Sprintf("Name: \"%s\",", r.Name),
		fmt.Sprintf("Type: \"%s\",", r.Type),
	}
	if r.Source != nil {
		parts = append(parts, fmt.Sprintf("Source: %#v,", r.Source))
	}
	if r.Privileged {
		parts = append(parts, fmt.Sprintf("Privileged: true,"))
	}
	if r.Params != nil {
		parts = append(parts, fmt.Sprintf("Params: %#v,", r.Params))
	}
	if r.CheckEvery != "" {
		parts = append(parts, fmt.Sprintf("CheckEvery: \"%s\",", r.CheckEvery))
	}
	if r.Tags != nil {
		parts = append(parts, fmt.Sprintf("Tags: %#v,", r.Tags))
	}

	// closing
	parts = append(parts, "}")

	name := fmt.Sprintf("ResourceType%s", r.Name)
	parts[0] = fmt.Sprintf("var %s = ResourceType{", name)

	generated := strings.Join(parts, "\n")

	NameToBlock[name] = generated

	return name
}
