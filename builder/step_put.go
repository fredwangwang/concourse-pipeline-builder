package builder

type StepPut struct {
	Put       string                 `yaml:"put,omitempty"`
	Resource  string                 `yaml:"resource,omitempty"`
	Params    map[string]interface{} `yaml:"params,omitempty"`
	GetParams map[string]interface{} `yaml:"get_params,omitempty"`
	StepHook  `yaml:",inline,omitempty"`
}

type _stepPut struct {
	Put       string                 `yaml:"put,omitempty"`
	Resource  string                 `yaml:"resource,omitempty"`
	Params    map[string]interface{} `yaml:"params,omitempty"`
	GetParams map[string]interface{} `yaml:"get_params,omitempty"`
}

func (s StepPut) StepType() string {
	return "put"
}

func (s *StepPut) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshalStep(s, &_stepPut{}, unmarshal)
}
