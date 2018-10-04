package builder

import (
	"fmt"
	"github.com/mitchellh/hashstructure"
	"log"
	"strings"
)

type StepGet struct {
	Get      string                 `yaml:"get,omitempty"`
	Resource string                 `yaml:"resource,omitempty"`
	Version  interface{}            `yaml:"version,omitempty"`
	Passed   []string               `yaml:"passed,omitempty"` //TODO: change to []Job?
	Params   map[string]interface{} `yaml:"params,omitempty"`
	Trigger  bool                   `yaml:"trigger,omitempty"`
	StepHook `yaml:",inline,omitempty"`
}

type _stepGet struct {
	Get      string                 `yaml:"get,omitempty"`
	Resource string                 `yaml:"resource,omitempty"`
	Version  interface{}            `yaml:"version,omitempty"`
	Passed   []string               `yaml:"passed,omitempty"` //TODO: change to []Job?
	Params   map[string]interface{} `yaml:"params,omitempty"`
	Trigger  bool                   `yaml:"trigger,omitempty"`
}

func (s StepGet) StepType() string {
	return "get"
}

func (s *StepGet) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshalStep(s, &_stepGet{}, unmarshal)
}

func (s StepGet) Generate() string {
	var parts = []string{
		"StepGet{", // placeholder
		fmt.Sprintf("Get: \"%s\",", s.Get),
	}
	if s.Resource != "" {
		parts = append(parts, fmt.Sprintf("Resource: \"%s\",", s.Resource))
	}
	if s.Version != nil {
		// as version can be ("every" | "latest" | obj), need to test what type is version
		if versionStr, ok := s.Version.(string); ok {
			parts = append(parts, fmt.Sprintf("Version: \"%s\",", versionStr))
		} else {
			parts = append(parts, fmt.Sprintf("Version: %#v,", s.Version))
		}
	}
	if s.Passed != nil {
		parts = append(parts, fmt.Sprintf("Passed: %#v,", s.Passed))
	}
	if s.Params != nil {
		parts = append(parts, fmt.Sprintf("Params: %#v,", s.Params))
	}
	if s.Trigger {
		parts = append(parts, fmt.Sprintf("Trigger: true,"))
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

	name := fmt.Sprintf("StepGet%s%x", sanitizeVarName(s.Get), hash)
	parts[0] = fmt.Sprintf("var %s = StepGet{", name)

	generated := strings.Join(parts, "\n")

	StepNameToBlock.Set(name, generated)

	return name
}
