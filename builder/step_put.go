package builder

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

//func (s *StepPut) Generate() string {
//	tmpl := `
//var StepPut%s = StepPut{
//	Put:       "%s",
//	Resource:  "%s",
//	Params:    %#v,
//	GetParams: %#v,
//    StepHook:  StepHook{
//        OnSuccess: %s,
//        OnFailure: %s,
//        OnAbort: %s,
//        Ensure: %s,
//        Tags: %#v,
//        Timeout: "%s",
//        Attempts: %d,
//    }
//}
//`
//
//	generated := fmt.Sprintf(tmpl,
//		"name", s.Put, s.Resource, s.Params, s.GetParams,
//		"a", "b", "c", "d", s.Tags, s.Timeout, s.Attempts)
//
//	//h := fnv.New32a()
//	//h.Write([]byte(generated))
//	//index := h.Sum32()
//	//
//	//IndexBlock[index] = generated
//	//
//	//fmt.Println(generated)
//
//	fmt.Println(generated)
//	s.OnSuccess = StepGet{
//		Get:"adsf",
//	}
//
//	fmt.Printf("%#v\n", s)
//
//	return ""
//}
