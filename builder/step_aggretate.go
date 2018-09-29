package builder

type StepAggregate struct {
	Aggregate Steps `yaml:"aggregate,omitempty"`
	StepHook  `yaml:",inline,omitempty"`
}

type _stepAggregate struct {
	Aggregate Steps `yaml:"aggregate,omitempty"`
}

func (s StepAggregate) StepType() string {
	return "aggregate"
}

func (s *StepAggregate) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshalStep(s, &_stepAggregate{}, unmarshal)
}
