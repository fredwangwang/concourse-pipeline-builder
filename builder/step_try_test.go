package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StepTry", func() {
	It("generates proper code section", func() {
		step1 := StepTry{
			Try: StepGet{
				Get: "try-get",
			},
			StepHook: StepHook{
				OnSuccess: StepDo{
					Do: Steps{
						StepGet{
							Get: "something",
						},
						StepPut{
							Put: "something",
						},
					},
				},
				OnFailure: StepDo{
					Do: Steps{
						StepGet{
							Get: "something",
						},
						StepPut{
							Put: "something",
						},
					},
				},
			},
		}

		expected := `var StepTry518229490 = StepTry{
Try: StepGettry-get3956449771,
StepHook:  StepHook{
OnSuccess: StepDo4211289418,
OnFailure: StepDo4211289418,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock[stepName]
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(Equal(expected))
	})
})
