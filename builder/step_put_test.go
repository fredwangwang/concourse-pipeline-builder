package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StepPut", func() {
	It("generates proper code section", func() {
		step1 := StepPut{
			Put:      "something",
			Resource: "some-res",
			Params: map[string]interface{}{
				"abc": "def",
			},
			GetParams: map[string]interface{}{
				"hello": 1234,
			},
			StepHook: StepHook{
				OnSuccess: StepPut{
					Put: "another",
				},
				OnFailure: StepPut{
					Put: "another",
				},
			},
		}

		expected := `var StepPutsomething439ab7eb3de6c2c = StepPut{
Put: "something",
Resource: "some-res",
Params: map[string]interface {}{"abc":"def"},
GetParams: map[string]interface {}{"hello":1234},
StepHook:  StepHook{
OnSuccess: StepPutanotheraa6dbabbb5ec4b64,
OnFailure: StepPutanotheraa6dbabbb5ec4b64,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock.Get(stepName)
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result.(string)))
		Expect(result).To(Equal(expected))
	})
})
