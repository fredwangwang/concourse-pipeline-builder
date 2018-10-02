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

		expected := `var StepDo3206072831 = StepDo{
Do: Steps{
StepGetsomething3739314951,
StepPutsomething1117426483,
},
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
