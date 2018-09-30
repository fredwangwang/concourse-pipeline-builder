package builder_test

import (
	. "github.com/onsi/ginkgo"
	"gopkg.in/yaml.v2"

	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/gomega"
)

var _ = Describe("Step", func() {
	var yamlStr = `
jobs:
- name: sample
  plan:
  - get: some-get
    trigger: true
    on_success:
      get: on-success
    on_failure:
      get: on-failure
    on_abort:
      get: on-abort
    ensure:
      get: ensure
    tags:
    - test
    timeout: 10m
    attempts: 10
  - put: some-put
    on_success:
      get: on-success
  - try:
      get: try-get
    on_success:
      get: on-success
  - do:
    - get: do-get1
    - get: do-get2
    on_success:
      get: on-success
  - aggregate:
    - get: aggregate-get1
    - get: aggregate-get2
    on_success:
      get: on-success
  - task: task-one
    config:
      platform: linux
      image_resource:
        type: docker-image
        source:
          repository: ubuntu
      params:
        SOME-FIELD: ~
      run:
        path: bash
        args: [ "\necho \"$SOME-FIELD\"\n" ]
    input_mapping:
      renamed-input: some-input
    params:
      SOME-FIELD: hello-world
    on_success:
      get: on-success
`

	var pipeStruct = Pipeline{
		Jobs: []Job{
			{
				Name: "sample",
				Plan: []Step{
					StepGet{
						Get:     "some-get",
						Trigger: true,
						StepHook: StepHook{
							OnSuccess: StepGet{
								Get: "on-success",
							},
							OnFailure: StepGet{
								Get: "on-failure",
							},
							OnAbort: StepGet{
								Get: "on-abort",
							},
							Ensure: StepGet{
								Get: "ensure",
							},
							Tags: []string{
								"test",
							},
							Timeout:  "10m",
							Attempts: 10,
						},
					},
					StepPut{
						Put: "some-put",
						StepHook: StepHook{
							OnSuccess: StepGet{
								Get: "on-success",
							},
						},
					},
					StepTry{
						Try: StepGet{
							Get: "try-get",
						},
						StepHook: StepHook{
							OnSuccess: StepGet{
								Get: "on-success",
							},
						},
					},
					StepDo{
						Do: Steps{
							StepGet{
								Get: "do-get1",
							},
							StepGet{
								Get: "do-get2",
							},
						},
						StepHook: StepHook{
							OnSuccess: StepGet{
								Get: "on-success",
							},
						},
					},
					StepAggregate{
						Aggregate: Steps{
							StepGet{
								Get: "aggregate-get1",
							},
							StepGet{
								Get: "aggregate-get2",
							},
						},
						StepHook: StepHook{
							OnSuccess: StepGet{
								Get: "on-success",
							},
						},
					},
					StepTask{
						Task: "task-one",
						Config: TaskConfig{
							Platform: "linux",
							ImageResource: TaskImageResource{
								Type: "docker-image",
								Source: map[string]interface{}{
									"repository": "ubuntu",
								},
							},
							Run: TaskRun{
								Path: "bash",
								Args: []string{
									`
echo "$SOME-FIELD"
`,
								},
							},
							Params: map[string]interface{}{
								"SOME-FIELD": nil,
							},
						},
						Params: map[string]interface{}{
							"SOME-FIELD": "hello-world",
						},
						InputMapping: map[string]interface{}{
							"renamed-input": "some-input",
						},
						StepHook: StepHook{
							OnSuccess: StepGet{
								Get: "on-success",
							},
						},
					},
				},
			},
		},
	}
	It("unmarshals step section", func() {
		var pipe Pipeline
		err := yaml.Unmarshal([]byte(yamlStr), &pipe)
		Expect(err).NotTo(HaveOccurred())
		Expect(pipe).To(Equal(pipeStruct))
	})

	It("marshals step section", func() {
		yamlBytes, err := yaml.Marshal(pipeStruct)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(yamlBytes)).To(MatchYAML(yamlStr))
	})
})
