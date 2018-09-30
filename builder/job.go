package builder

type Job struct {
	Name                 string   `yaml:"name,omitempty"`
	Plan                 Steps    `yaml:"plan,omitempty"`
	Serial               bool     `yaml:"serial,omitempty"`
	SerialGroups         []string `yaml:"serial_groups,omitempty"`
	BuildLogsToRetain    int      `yaml:"build_logs_to_retain,omitempty"`
	MaxInFlight          int      `yaml:"max_in_flight,omitempty"` // TODO: make default 1
	Public               bool     `yaml:"public,omitempty"`
	DisableManualTrigger bool     `yaml:"disable_manual_trigger,omitempty"`
	Interruptible        bool     `yaml:"interruptible,omitempty"`
	OnSuccess            Step     `yaml:"on_success,omitempty"`
	OnFailure            Step     `yaml:"on_failure,omitempty"`
	OnAbort              Step     `yaml:"on_abort,omitempty"`
	Ensure               Step     `yaml:"ensure,omitempty"`
}

type _job struct {
	Name                 string   `yaml:"name,omitempty"`
	Plan                 Steps    `yaml:"plan,omitempty"`
	Serial               bool     `yaml:"serial,omitempty"`
	SerialGroups         []string `yaml:"serial_groups,omitempty"`
	BuildLogsToRetain    int      `yaml:"build_logs_to_retain,omitempty"`
	MaxInFlight          int      `yaml:"max_in_flight,omitempty"`
	Public               bool     `yaml:"public,omitempty"`
	DisableManualTrigger bool     `yaml:"disable_manual_trigger,omitempty"`
	Interruptible        bool     `yaml:"interruptible,omitempty"`
}

func (j *Job) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshalStep(j, &_job{}, unmarshal)
}
