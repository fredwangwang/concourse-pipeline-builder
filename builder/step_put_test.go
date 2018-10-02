package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StepPut", func() {
	It("generates proper code section", func() {
		step1 := StepPut{
			Put:       "something",
			Resource:  "some-res",
			Params:    nil,
			GetParams: nil,
			StepHook: StepHook{
				OnSuccess: StepPut{
					Put: "another",
				},
				OnFailure: StepPut{
					Put: "another",
				},
			},
		}

		expected := `var StepPutsomething898161121 = StepPut{
Put: "something",
Resource: "some-res",
StepHook:  StepHook{
OnSuccess: "StepPutanother1461752950",
OnFailure: "StepPutanother1461752950",
},
}`

		stepName := step1.Generate()
		Expect(StepNameToBlock[stepName]).To(Equal(expected))
	})
})
