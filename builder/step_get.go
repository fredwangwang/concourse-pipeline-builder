package builder

type StepGet struct {
	Get      string                 `yaml:"get,omitempty"`
	Resource string                 `yaml:"resource,omitempty"`
	Version  interface{}            `yaml:"version,omitempty"`
	Passed   []string               `yaml:"passed,omitempty"` //TODO: change to []Job?
	Params   map[string]interface{} `yaml:"params,omitempty"`
	Trigger  bool                   `yaml:"trigger,omitempty"`
	StepHook `yaml:",inline,omitempty"`
}

type _stepGet struct {
	Get      string                 `yaml:"get,omitempty"`
	Resource string                 `yaml:"resource,omitempty"`
	Version  interface{}            `yaml:"version,omitempty"`
	Passed   []string               `yaml:"passed,omitempty"` //TODO: change to []Job?
	Params   map[string]interface{} `yaml:"params,omitempty"`
	Trigger  bool                   `yaml:"trigger,omitempty"`
}

func (s StepGet) StepType() string {
	return "get"
}

func (s *StepGet) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshalStep(s, &_stepGet{}, unmarshal)
}
