package builder

type StepDo struct {
	Do       Steps `yaml:"do,omitempty"`
	StepHook `yaml:",inline,omitempty"`
}

type _stepDo struct {
	Do Steps `yaml:"do,omitempty"`
}

func (s StepDo) StepType() string {
	return "do"
}

func (s *StepDo) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshalStep(s, &_stepDo{}, unmarshal)
}

func (s StepDo) Generate() string {
	return ""
}
