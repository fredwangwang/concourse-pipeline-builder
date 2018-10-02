package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StepGet", func() {
	It("generates proper code section", func() {
		step1 := StepGet{
			Get:      "something",
			Resource: "some-res",
			Version: map[string]string{
				"ref": "adsfjklajdfkl",
			},
			Passed:  []string{"a", "b"},
			Trigger: true,
			Params: map[string]interface{}{
				"abc": "def",
			},
			StepHook: StepHook{
				OnSuccess: StepGet{
					Get: "another",
				},
				OnFailure: StepGet{
					Get: "another",
				},
			},
		}

		expected := `var StepGetsomething817928348 = StepGet{
Get: "something",
Resource: "some-res",
Version: map[string]string{"ref":"adsfjklajdfkl"},
Passed: []string{"a", "b"},
Params: map[string]interface {}{"abc":"def"},
Trigger: true,
StepHook:  StepHook{
OnSuccess: StepGetanother2093271898,
OnFailure: StepGetanother2093271898,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock[stepName]
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(Equal(expected))
	})
})
