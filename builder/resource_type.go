package builder

import (
	"gopkg.in/go-playground/validator.v9"
)

// this interface does not do anything special now, but serve mainly as a schema enforcement.
type ResourceType interface {
	ResourceType() string
	Validate() error
}

type ResourceTypes []ResourceType

type ResourceTypeGeneric struct {
	Name       string                 `yaml:"name,omitempty" validate:"required"`
	Type       string                 `yaml:"type,omitempty" validate:"required"`
	Source     map[string]interface{} `yaml:"source,omitempty"`
	Privileged bool                   `yaml:"privileged,omitempty"`
	Params     map[string]interface{} `yaml:"params,omitempty"`
	CheckEvery string                 `yaml:"check_every,omitempty"`
	Tags       []string               `yaml:"tags,omitempty"`
}

func (r *ResourceTypes) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// TODO: support types other than generic.
	var holder []ResourceTypeGeneric

	err := unmarshal(&holder)
	if err != nil {
		return err
	}

	for _, item := range holder {
		*r = append(*r, item)
	}

	return nil
}

func (r ResourceTypeGeneric) ResourceType() string {
	return r.Type
}

func (r ResourceTypeGeneric) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
