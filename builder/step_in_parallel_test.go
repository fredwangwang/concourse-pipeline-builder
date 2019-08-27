package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("StepInParallel", func() {
	var step1 = StepInParallel{
		InParallel: InParallel{
			Limit:    2,
			FailFast: true,
			Steps: Steps{
				StepGet{
					Get: "something",
				},
				StepPut{
					Put: "something",
				},
			},
		},
		StepHook: StepHook{
			OnSuccess: StepInParallel{
				InParallel: InParallel{
					Steps: Steps{
						StepGet{
							Get: "something",
						},
						StepPut{
							Put: "something",
						},
					},
				},
			},
			OnFailure: StepInParallel{
				InParallel: InParallel{
					Steps: Steps{
						StepGet{
							Get: "something",
						},
						StepPut{
							Put: "something",
						},
					},
				},
			},
		},
	}

	It("generates proper code section", func() {
		expected := `var StepInParallelc61de7380e177b40 = StepInParallel{
InParallel: InParallel{
Limit: 2,
FailFast: true,
Steps: Steps{
StepGetsomething616bd9df0a81a013,
StepPutsomethingcc7ec3f5209f51db,
},
},
StepHook:  StepHook{
OnSuccess: StepInParallelea9bf7a7cb7473c0,
OnFailure: StepInParallelea9bf7a7cb7473c0,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock.Get(stepName)
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result.(string)))
		Expect(result).To(Equal(expected))
	})

	It("marshals", func() {
		expected := `---
in_parallel:
  fail_fast: true
  limit: 2
  steps:
  - get: something
  - put: something
on_failure:
  in_parallel:
    steps:
    - get: something
    - put: something
on_success:
  in_parallel:
    steps:
    - get: something
    - put: something
`

		contentBytes, err := yaml.Marshal(step1)
		Expect(err).NotTo(HaveOccurred())
		Expect(expected).To(MatchYAML(contentBytes))
	})
})
