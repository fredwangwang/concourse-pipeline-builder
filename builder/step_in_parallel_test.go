package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StepInParallel", func() {
	It("generates proper code section", func() {
		step1 := StepInParallel{
			InParallel: Steps{
				StepGet{
					Get: "something",
				},
				StepPut{
					Put: "something",
				},
			},
			StepHook: StepHook{
				OnSuccess: StepInParallel{
					InParallel: Steps{
						StepGet{
							Get: "something",
						},
						StepPut{
							Put: "something",
						},
					},
				},
				OnFailure: StepInParallel{
					InParallel: Steps{
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

		expected := `var StepInParalleleb15d93e38b0d287 = StepInParallel{
InParallel: Steps{
StepGetsomething616bd9df0a81a013,
StepPutsomethingcc7ec3f5209f51db,
},
StepHook:  StepHook{
OnSuccess: StepInParallel42373b713e82190f,
OnFailure: StepInParallel42373b713e82190f,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock.Get(stepName)
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result.(string)))
		Expect(result).To(Equal(expected))
	})
})
