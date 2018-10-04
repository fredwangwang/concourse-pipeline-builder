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

		expected := `var StepTasksomethingf7b0884af87efc8f = StepTask{
Task: "something",
File: "task.yml",
Privileged: true,
Params: map[string]interface {}{"abc":"def"},
Image: "abc",
InputMapping: map[string]interface {}{"in":"somehting"},
StepHook:  StepHook{
OnSuccess: StepPutanotheraa6dbabbb5ec4b64,
OnFailure: StepPutanotheraa6dbabbb5ec4b64,
},
}`

		stepName := step1.Generate()
		result, ok := StepNameToBlock.Get(stepName)
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result.(string)))
		Expect(result).To(Equal(expected))
	})
})
