package builder

import (
	"fmt"
	"reflect"
	"time"
)

type Step interface {
	StepType() string
	Generate() string
}

type Steps []Step

func (s *Steps) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var holder []map[string]interface{}

	err := unmarshal(&holder)
	if err != nil {
		return err
	}

	for _, stepMap := range holder {
		stepStruct, err := parseStep(stepMap)
		if err != nil {
			return err
		}
		*s = append(*s, stepStruct)
	}

	return nil
}

// TODO: add validation for all step structs
type StepHook struct {
	OnSuccess Step     `yaml:"on_success,omitempty"`
	OnFailure Step     `yaml:"on_failure,omitempty"`
	OnAbort   Step     `yaml:"on_abort,omitempty"`
	Ensure    Step     `yaml:"ensure,omitempty"`
	Tags      []string `yaml:"tags,omitempty"`
	Timeout   string   `yaml:"timeout,omitempty"` // TODO: validate time duration
	Attempts  int      `yaml:"attempts,omitempty"`
}

type _stepHook struct {
	OnSuccess map[string]interface{} `yaml:"on_success,omitempty"`
	OnFailure map[string]interface{} `yaml:"on_failure,omitempty"`
	OnAbort   map[string]interface{} `yaml:"on_abort,omitempty"`
	Ensure    map[string]interface{} `yaml:"ensure,omitempty"`
	Tags      []string               `yaml:"tags,omitempty"`
	Timeout   string                 `yaml:"timeout,omitempty"`
	Attempts  int                    `yaml:"attempts,omitempty"`
}

// Due to the implementation of yaml:",inline", the inlined struct's field is processed
// as if they are the fields of the outer struct. Which means the custom unmarshal rule
// of the inlined StepHook struct will not be applied automatically.
// Thus when unmarshaling steps, the following unmarshal rule needs to be applied manually
// in each step's custom unmarshal rule.
func (s *StepHook) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var internal _stepHook

	err := unmarshal(&internal)
	if err != nil {
		return err
	}

	if internal.Timeout != "" {
		_, err = time.ParseDuration(internal.Timeout)
		if err != nil {
			return err
		}
	}

	s.Tags = internal.Tags
	s.Timeout = internal.Timeout
	s.Attempts = internal.Attempts

	s.OnSuccess, err = parseStep(internal.OnSuccess)
	if err != nil {
		return err
	}
	s.OnFailure, err = parseStep(internal.OnFailure)
	if err != nil {
		return err
	}
	s.OnAbort, err = parseStep(internal.OnAbort)
	if err != nil {
		return err
	}
	s.Ensure, err = parseStep(internal.Ensure)
	if err != nil {
		return err
	}

	return nil
}

func parseStep(stepMap map[string]interface{}) (Step, error) {
	var err error
	var res Step

	if stepMap == nil || len(stepMap) == 0 {
		return nil, nil
	}

	if _, ok := stepMap["get"]; ok {
		stepStruct := StepGet{}
		err = yamlTransform(stepMap, &stepStruct)
		if err != nil {
			return nil, err
		}
		res = stepStruct
	} else if _, ok := stepMap["put"]; ok {
		stepStruct := StepPut{}
		err = yamlTransform(stepMap, &stepStruct)
		if err != nil {
			return nil, err
		}
		res = stepStruct
	} else if _, ok := stepMap["task"]; ok {
		stepStruct := StepTask{}
		err = yamlTransform(stepMap, &stepStruct)
		if err != nil {
			return nil, err
		}
		res = stepStruct
	} else if _, ok := stepMap["aggregate"]; ok {
		stepStruct := StepAggregate{}
		err = yamlTransform(stepMap, &stepStruct)
		if err != nil {
			return nil, err
		}
		res = stepStruct
	} else if _, ok := stepMap["in_parallel"]; ok {
		stepStruct := StepInParallel{}
		err = yamlTransform(stepMap, &stepStruct)
		if err != nil {
			return nil, err
		}
		res = stepStruct
	} else if _, ok := stepMap["do"]; ok {
		stepStruct := StepDo{}
		err = yamlTransform(stepMap, &stepStruct)
		if err != nil {
			return nil, err
		}
		res = stepStruct
	} else if _, ok := stepMap["try"]; ok {
		stepStruct := StepTry{}
		err = yamlTransform(stepMap, &stepStruct)
		if err != nil {
			return nil, err
		}
		res = stepStruct
	} else {
		// unrecognized step
		return nil, fmt.Errorf("unrecognized step: %+v", stepMap)
	}
	return res, nil
}

func unmarshalStep(dst interface{}, internal interface{}, unmarshal func(interface{}) error) error {
	var stepHook StepHook

	err := unmarshal(&stepHook)
	if err != nil {
		return err
	}

	err = unmarshal(internal)
	if err != nil {
		return err
	}

	typeInternal := reflect.TypeOf(internal).Elem()
	valInternal := reflect.ValueOf(internal).Elem()

	typeStep := reflect.TypeOf(dst).Elem()
	valStep := reflect.ValueOf(dst).Elem()

	for i := 0; i < typeInternal.NumField(); i++ {
		fieldName := typeInternal.Field(i).Name
		stepFieldType, _ := typeStep.FieldByName(fieldName)
		if stepFieldType.Type != typeInternal.Field(i).Type {
			// try convert valInternal.Field(i) to map[string]interface{},
			// as that is the type being used for all intermediate Step unmarshal.
			// The reason to have a intermediate map[string]interface{} is because
			// there is no way to direct unmarshal the data into an interface.
			// So additional conversion is required to perform the proper data unmarshalling.

			// no need to check, if it is not the type expected, the program really should just
			// stop since that is not expected at all.
			internalFieldVal := valInternal.Field(i).Interface().(map[string]interface{})

			stepFieldVal, err := parseStep(internalFieldVal)
			if err != nil {
				return err
			}

			valStep.FieldByName(fieldName).Set(reflect.ValueOf(stepFieldVal))
		} else {
			valStep.FieldByName(fieldName).Set(valInternal.Field(i))
		}
	}

	typeStepHook := reflect.TypeOf(stepHook)
	valStepHook := reflect.ValueOf(stepHook)

	for i := 0; i < typeStepHook.NumField(); i++ {
		fieldName := typeStepHook.Field(i).Name
		if _, ok := typeStep.FieldByName(fieldName); !ok {
			continue
		}
		valStep.FieldByName(fieldName).Set(valStepHook.Field(i))
	}

	return nil
}
