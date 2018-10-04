package builder

import (
	"fmt"
	"github.com/mitchellh/hashstructure"
	"log"
	"strings"
)

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

func (s StepAggregate) Generate() string {
	var parts = []string{
		"StepGet{", // placeholder
	}
	if s.Aggregate != nil {
		parts = append(parts, "Aggregate: Steps{")

		for _, step := range s.Aggregate {
			stepName := step.Generate()
			parts = append(parts, fmt.Sprintf("%s,", stepName))
		}
		parts = append(parts, "},")
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

	// do not see a particular reason why this would ever fail. If it fails, let it fail then.
	hash, err := hashstructure.Hash(s, nil)
	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("StepAggregate%x", hash)
	parts[0] = fmt.Sprintf("var %s = StepAggregate{", name)

	generated := strings.Join(parts, "\n")

	StepNameToBlock.Set(name, generated)

	return name
}
