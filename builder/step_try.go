package builder

type StepTry struct {
	Try      Step `yaml:"try,omitempty"`
	StepHook `yaml:",inline,omitempty"`
}

type _stepTry struct {
	Try map[string]interface{} `yaml:"try,omitempty"`
}

func (s StepTry) StepType() string {
	return "try"
}

func (s *StepTry) UnmarshalYAML(unmarshal func(interface{}) error) error {
	unmarshalStep(s, &_stepTry{}, unmarshal)
	return nil
}

func (s StepTry) Generate() string {
	return ""
}
