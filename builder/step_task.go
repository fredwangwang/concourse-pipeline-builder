package builder

type StepTask struct {
	Task          string                 `yaml:"task,omitempty"`
	Config        TaskConfig             `yaml:"config,omitempty"` // TODO: validate config and file can have only one set
	File          string                 `yaml:"file,omitempty"`
	Privileged    bool                   `yaml:"privileged,omitempty"`
	Params        map[string]interface{} `yaml:"params,omitempty"`
	Image         string                 `yaml:"image,omitempty"`
	InputMapping  map[string]interface{} `yaml:"input_mapping,omitempty"`
	OutputMapping map[string]interface{} `yaml:"output_mapping,omitempty"`
	StepHook      `yaml:",inline,omitempty"`
}

type _stepTask struct {
	Task          string                 `yaml:"task,omitempty"`
	Config        TaskConfig             `yaml:"config,omitempty"` // TODO: validate config and file can have only one set
	File          string                 `yaml:"file,omitempty"`
	Privileged    bool                   `yaml:"privileged,omitempty"`
	Params        map[string]interface{} `yaml:"params,omitempty"`
	Image         string                 `yaml:"image,omitempty"`
	InputMapping  map[string]interface{} `yaml:"input_mapping,omitempty"`
	OutputMapping map[string]interface{} `yaml:"output_mapping,omitempty"`
}

func (s StepTask) StepType() string {
	return "task"
}

func (s *StepTask) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshalStep(s, &_stepTask{}, unmarshal)
}
