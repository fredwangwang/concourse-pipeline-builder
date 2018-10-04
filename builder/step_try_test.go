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

		expected := `var StepTryc05e731c896b6eed = StepTry{
Try: StepGettryget1b5dfff4bf2ce5e,
StepHook:  StepHook{
OnSuccess: StepDoefcaa20bd1df0b03,
OnFailure: StepDoefcaa20bd1df0b03,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock.Get(stepName)
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result.(string)))
		Expect(result).To(Equal(expected))
	})
})
