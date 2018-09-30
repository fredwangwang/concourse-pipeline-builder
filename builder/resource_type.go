package builder

import (
	"gopkg.in/go-playground/validator.v9"
)

type ResourceType struct {
	Name       string                 `yaml:"name,omitempty" validate:"required"`
	Type       string                 `yaml:"type,omitempty" validate:"required"`
	Source     map[string]interface{} `yaml:"source,omitempty"`
	Privileged bool                   `yaml:"privileged,omitempty"`
	Params     map[string]interface{} `yaml:"params,omitempty"`
	CheckEvery string                 `yaml:"check_every,omitempty"`
	Tags       []string               `yaml:"tags,omitempty"`
}

type ResourceTypes []ResourceType

//func (r *ResourceTypes) UnmarshalYAML(unmarshal func(interface{}) error) error {
//	var holder []ResourceType
//
//	err := unmarshal(&holder)
//	if err != nil {
//		return err
//	}
//
//	for _, item := range holder {
//		*r = append(*r, item)
//	}
//
//	return nil
//}

func (r ResourceType) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
