package builder

import (
	"fmt"
	"github.com/mitchellh/hashstructure"
	"log"
	"strings"
)

type InParallel struct {
	Limit    int   `yaml:"limit,omitempty"`
	FailFast bool  `yaml:"fail_fast,omitempty"`
	Steps    Steps `yaml:"steps,omitempty"`
}

type StepInParallel struct {
	InParallel InParallel `yaml:"in_parallel,omitempty"`
	StepHook   `yaml:",inline,omitempty"`
}

type _stepInParallel struct {
	InParallel InParallel `yaml:"in_parallel,omitempty"`
}

func (s StepInParallel) StepType() string {
	return "in_parallel"
}

func (s *StepInParallel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshalStep(s, &_stepInParallel{}, unmarshal)
}

func (s StepInParallel) Generate() string {
	var parts = []string{
		"", // placeholder
	}

	parts = append(parts, "InParallel: InParallel{")
	if s.InParallel.Limit != 0 {
		parts = append(parts, fmt.Sprintf("Limit: %d,", s.InParallel.Limit))
	}
	if s.InParallel.FailFast {
		parts = append(parts, "FailFast: true,")
	}
	parts = append(parts, "Steps: Steps{")
	for _, step := range s.InParallel.Steps {
		stepName := step.Generate()
		parts = append(parts, fmt.Sprintf("%s,", stepName))
	}
	parts = append(parts, "},")
	parts = append(parts, "},")

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

	// do not see a particular reason why this would ever fail. If it fails, let it fail then.
	hash, err := hashstructure.Hash(s, nil)
	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("StepInParallel%x", hash)
	parts[0] = fmt.Sprintf("var %s = StepInParallel{", name)

	generated := strings.Join(parts, "\n")

	StepNameToBlock.Set(name, generated)

	return name
}
