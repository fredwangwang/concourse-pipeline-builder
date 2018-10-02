package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StepAggregate", func() {
	It("generates proper code section", func() {
		step1 := StepAggregate{
			Aggregate: Steps{
				StepGet{
					Get: "something",
				},
				StepPut{
					Put: "something",
				},
			},
			StepHook: StepHook{
				OnSuccess: StepAggregate{
					Aggregate: Steps{
						StepGet{
							Get: "something",
						},
						StepPut{
							Put: "something",
						},
					},
				},
				OnFailure: StepAggregate{
					Aggregate: Steps{
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

		expected := `var StepAggregatefcea162f7e2c5285 = StepAggregate{
Aggregate: Steps{
StepGetsomething616bd9df0a81a013,
StepPutsomethingcc7ec3f5209f51db,
},
StepHook:  StepHook{
OnSuccess: StepAggregate6ec501561a44889f,
OnFailure: StepAggregate6ec501561a44889f,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock[stepName]
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(Equal(expected))
	})
})
