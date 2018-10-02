package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StepTask", func() {
	It("generates proper code section", func() {
		step1 := StepTask{
			Task:       "something",
			Config:     nil,
			File:       "task.yml",
			Privileged: true,
			Params: map[string]interface{}{
				"abc": "def",
			},
			Image: "abc",
			InputMapping: map[string]interface{}{
				"in": "somehting",
			},
			OutputMapping: nil,
			StepHook: StepHook{
				OnSuccess: StepPut{
					Put: "another",
				},
				OnFailure: StepPut{
					Put: "another",
				},
			},
		}

		expected := `var StepTasksomething3417063671 = StepTask{
Task: "something",
File: "task.yml",
Privileged: true,
Params: map[string]interface {}{"abc":"def"},
Image: "abc",
InputMapping: map[string]interface {}{"in":"somehting"},
StepHook:  StepHook{
OnSuccess: StepPutanother1461752950,
OnFailure: StepPutanother1461752950,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock[stepName]
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(Equal(expected))
	})
})
