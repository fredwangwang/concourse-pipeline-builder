package builder

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
	return ""
}