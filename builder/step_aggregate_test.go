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

		expected := `var StepAggregate2201624517 = StepAggregate{
Aggregate: Steps{
StepGetsomething3739314951,
StepPutsomething1117426483,
},
StepHook:  StepHook{
OnSuccess: StepAggregate238783884,
OnFailure: StepAggregate238783884,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock[stepName]
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(Equal(expected))
	})
})
