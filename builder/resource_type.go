package builder

type ResourceType struct {
	Name       string                 `yaml:"name,omitempty"`
	Type       string                 `yaml:"type,omitempty"`
	Source     map[string]interface{} `yaml:"source,omitempty"`
	Privileged bool                   `yaml:"privileged,omitempty"`
	Params     map[string]interface{} `yaml:"params,omitempty"`
	CheckEvery string                 `yaml:"check_every,omitempty"`
	Tags       []string               `yaml:"tags,omitempty"`
}
