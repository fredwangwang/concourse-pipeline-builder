package builder

import (
	"fmt"
	"strings"
)

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

func (j Job) Generate() string {
	var parts = []string{
		"Job{", // placeholder
		fmt.Sprintf("Name: \"%s\",", j.Name),
	}
	if j.Plan != nil {
		parts = append(parts, "Plan: Steps{")
		for _, step := range j.Plan {
			parts = append(parts, fmt.Sprintf("%s,", step.Generate()))
		}
		parts = append(parts, "},")
	}
	if j.Serial {
		parts = append(parts, "Serial: true,")
	}
	if j.SerialGroups != nil {
		parts = append(parts, fmt.Sprintf("SerialGroups: %#v,", j.SerialGroups))
	}
	if j.BuildLogsToRetain != 0 {
		parts = append(parts, fmt.Sprintf("BuildLogsToRetain: %d,", j.BuildLogsToRetain))
	}
	if j.MaxInFlight != 0 {
		parts = append(parts, fmt.Sprintf("MaxInFlight: %d,", j.MaxInFlight))
	}
	if j.Public {
		parts = append(parts, "Public: true,")
	}
	if j.DisableManualTrigger {
		parts = append(parts, "DisableManualTrigger: true,")
	}
	if j.Interruptible {
		parts = append(parts, "Interruptible: true,")
	}
	if j.OnSuccess != nil {
		parts = append(parts, fmt.Sprintf("OnSuccess: %s,", j.OnSuccess.Generate()))
	}
	if j.OnFailure != nil {
		parts = append(parts, fmt.Sprintf("OnFailure: %s,", j.OnFailure.Generate()))
	}
	if j.OnAbort != nil {
		parts = append(parts, fmt.Sprintf("OnAbort: %s,", j.OnAbort.Generate()))
	}
	if j.Ensure != nil {
		parts = append(parts, fmt.Sprintf("Ensure: %s,", j.Ensure.Generate()))
	}

	// closing
	parts = append(parts, "}")

	name := fmt.Sprintf("Job%s", sanitizeVarName(j.Name))
	parts[0] = fmt.Sprintf("var %s = Job{", name)

	generated := strings.Join(parts, "\n")

	NameToBlock[name] = generated

	return name
}
