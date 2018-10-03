package builder

import (
	"fmt"
	"github.com/mitchellh/hashstructure"
	"log"
	"strings"
)

type StepPut struct {
	Put       string                 `yaml:"put,omitempty"`
	Resource  string                 `yaml:"resource,omitempty"`
	Params    map[string]interface{} `yaml:"params,omitempty"`
	GetParams map[string]interface{} `yaml:"get_params,omitempty"`
	StepHook  `yaml:",inline,omitempty"`
}

type _stepPut struct {
	Put       string                 `yaml:"put,omitempty"`
	Resource  string                 `yaml:"resource,omitempty"`
	Params    map[string]interface{} `yaml:"params,omitempty"`
	GetParams map[string]interface{} `yaml:"get_params,omitempty"`
}

func (s StepPut) StepType() string {
	return "put"
}

func (s *StepPut) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return unmarshalStep(s, &_stepPut{}, unmarshal)
}

func (s StepPut) Generate() string {
	var parts = []string{
		"StepPut{", // placeholder
		fmt.Sprintf("Put: \"%s\",", s.Put),
	}
	if s.Resource != "" {
		parts = append(parts, fmt.Sprintf("Resource: \"%s\",", s.Resource))
	}
	if s.Params != nil {
		parts = append(parts, fmt.Sprintf("Params: %#v,", s.Params))
	}
	if s.GetParams != nil {
		parts = append(parts, fmt.Sprintf("GetParams: %#v,", s.GetParams))
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
	// hash is deterministic. Given the same struct, the hash is always the same. // TODO: better desc
	// do not see a particular reason why this would ever fail. If it fails, let it fail then.
	hash, err := hashstructure.Hash(s, nil)
	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("StepPut%s%x", sanitizeVarName(s.Put), hash)
	parts[0] = fmt.Sprintf("var %s = StepPut{", name)

	generated := strings.Join(parts, "\n")

	NameToBlock[name] = generated

	return name
}
