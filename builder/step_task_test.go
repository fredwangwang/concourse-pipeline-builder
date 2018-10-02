package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = XDescribe("StepTask", func() {
	It("generates proper code section", func() {
		step1 := StepTask{
			Task:       "something",
			//Config:     nil,
			Params: map[string]interface{}{
				"abc": "def",
			},
			Image:         "abc",
			InputMapping: map[string]interface{}{
				"in": "somehting",
			},
			OutputMapping: nil,
			StepHook: StepHook{
				OnSuccess: StepPut{
					Put: "another",
				},
				OnFailure: StepPut{
					Put: "another",
				},
			},
		}

		expected := `var StepPutsomething3402357338 = StepPut{
Put: "something",
Resource: "some-res",
Params: map[string]interface {}{"abc":"def"},
GetParams: map[string]interface {}{"hello":1234},
StepHook:  StepHook{
OnSuccess: StepPutanother1461752950,
OnFailure: StepPutanother1461752950,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock[stepName]
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(Equal(expected))
	})
})
