package builder

import (
	"fmt"
	"github.com/mitchellh/hashstructure"
	"log"
	"strings"
)

type StepTask struct {
	Task          string                 `yaml:"task,omitempty"`
	Config        *TaskConfig            `yaml:"config,omitempty"` // TODO: validate config and file can have only one set
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
	Config        *TaskConfig            `yaml:"config,omitempty"` // TODO: validate config and file can have only one set
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

func (s StepTask) Generate() string {
	var parts = []string{
		"StepTask:{", // placeholder
		fmt.Sprintf("Task: \"%s\",", s.Task),
	}
	if s.Config != nil {
		parts = append(parts, fmt.Sprintf("Config: &%s,", s.Config.Generate()))
	}
	if s.File != "" {
		parts = append(parts, fmt.Sprintf("File: \"%s\",", s.File))
	}
	if s.Privileged {
		parts = append(parts, fmt.Sprintf("Privileged: true,"))
	}
	if s.Params != nil {
		parts = append(parts, fmt.Sprintf("Params: %#v,", s.Params))
	}
	if s.Image != "" {
		parts = append(parts, fmt.Sprintf("Image: \"%s\",", s.Image))
	}
	if s.InputMapping != nil {
		parts = append(parts, fmt.Sprintf("InputMapping: %#v,", s.InputMapping))
	}
	if s.OutputMapping != nil {
		parts = append(parts, fmt.Sprintf("OutputMapping: %#v,", s.OutputMapping))
	}

	// add stephook
	parts = append(parts, "StepHook:  StepHook{")
	if s.OnSuccess != nil {
		parts = append(parts, fmt.Sprintf("OnSuccess: %s,", s.OnSuccess.Generate()))
	}
	if s.OnFailure != nil {
		parts = append(parts, fmt.Sprintf("OnFailure: %s,", s.OnFailure.Generate()))
	}
	if s.OnAbort != nil {
		parts = append(parts, fmt.Sprintf("OnAbort: %s,", s.OnAbort.Generate()))
	}
	if s.Ensure != nil {
		parts = append(parts, fmt.Sprintf("Ensure: %s,", s.Ensure.Generate()))
	}
	if s.Tags != nil {
		parts = append(parts, fmt.Sprintf("Tags: %#v,", s.Tags))
	}
	if s.Timeout != "" {
		parts = append(parts, fmt.Sprintf("Timeout: \"%s\",", s.Timeout))
	}
	if s.Attempts != 0 {
		parts = append(parts, fmt.Sprintf("Attempts: %d,", s.Attempts))
	}
	parts = append(parts, "},")

	// closing
	parts = append(parts, "}")

	// get name
	// do not see a particular reason why this would ever fail. If it fails, let it fail then.
	hash, err := hashstructure.Hash(s, nil)
	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("StepTask%s%x", sanitizeVarName(s.Task), hash)
	parts[0] = fmt.Sprintf("var %s = StepTask{", name)

	generated := strings.Join(parts, "\n")

	NameToBlock[name] = generated

	return name
}
