package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StepDo", func() {
	It("generates proper code section", func() {
		step1 := StepDo{
			Do: Steps{
				StepGet{
					Get: "something",
				},
				StepPut{
					Put: "something",
				},
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

		expected := `var StepDo7529fcd10a2ae5ff = StepDo{
Do: Steps{
StepGetsomething616bd9df0a81a013,
StepPutsomethingcc7ec3f5209f51db,
},
StepHook:  StepHook{
OnSuccess: StepDoefcaa20bd1df0b03,
OnFailure: StepDoefcaa20bd1df0b03,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock[stepName]
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(Equal(expected))
	})
})
