package builder

type Step interface{}

type StepHook struct {
	OnSuccess Step     `yaml:"on_success,omitempty"`
	OnFailure Step     `yaml:"on_failure,omitempty"`
	OnAbort   Step     `yaml:"on_abort,omitempty"`
	Ensure    Step     `yaml:"ensure,omitempty"`
	Tags      []string `yaml:"tags,omitempty"`
	Timeout   string   `yaml:"timeout,omitempty"` // TODO: validate time duration
	Attempts  int      `yaml:"attempts,omitempty"`
}

type StepGet struct {
	Get      string                 `yaml:"get,omitempty"`
	Resource string                 `yaml:"resource,omitempty"`
	Version  interface{}            `yaml:"version,omitempty"`
	Passed   []string               `yaml:"passed,omitempty"`
	Params   map[string]interface{} `yaml:"params,omitempty"`
	Trigger  bool                   `yaml:"trigger,omitempty"`
	StepHook `yaml:",inline,omitempty"`
}

type StepPut struct {
	Put       string                 `yaml:"put,omitempty"`
	Resource  string                 `yaml:"resource,omitempty"`
	Params    map[string]interface{} `yaml:"params,omitempty"`
	GetParams map[string]interface{} `yaml:"get_params,omitempty"`
	StepHook  `yaml:",inline,omitempty"`
}

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

type StepAggregate struct {
	Aggregate []Step `yaml:"aggregate,omitempty"`
	StepHook  `yaml:",inline,omitempty"`
}

type StepDo struct {
	Do       []Step `yaml:"do,omitempty"`
	StepHook `yaml:",inline,omitempty"`
}

type StepTry struct {
	Try      Step `yaml:"try,omitempty"`
	StepHook `yaml:",inline,omitempty"`
}
