package builder_test

import (
	. "github.com/onsi/ginkgo"
	"gopkg.in/yaml.v2"

	. "github.com/onsi/gomega"

	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
)

var _ = Describe("Job", func() {
	var yamlStr = `
jobs:
- name: job1
  plan:
  - get: something
  serial: true
  serial_groups: [ 'a' ]
  build_logs_to_retain: 100
  max_in_flight: 1
  public: true
  disable_manual_trigger: true
  on_success:
    get: success
  on_failure:
    get: failure
  on_abort:
    get: abort
  ensure:
    get: ensure
- name: job2
  plan:
  - get: something-else
`

	var pipeStruct = Pipeline{
		Jobs: []Job{
			{
				Name: "job1",
				Plan: Steps{
					StepGet{
						Get: "something",
					},
				},
				Serial:               true,
				SerialGroups:         []string{"a"},
				BuildLogsToRetain:    100,
				MaxInFlight:          1,
				Public:               true,
				DisableManualTrigger: true,
				OnSuccess: StepGet{
					Get: "success",
				},
				OnFailure: StepGet{
					Get: "failure",
				},
				OnAbort: StepGet{
					Get: "abort",
				},
				Ensure: StepGet{
					Get: "ensure",
				},
			},
			{
				Name: "job2",
				Plan: Steps{
					StepGet{
						Get: "something-else",
					},
				},
			},
		},
	}

	It("unmarshals the job section", func() {
		var pipe Pipeline
		err := yaml.Unmarshal([]byte(yamlStr), &pipe)
		Expect(err).NotTo(HaveOccurred())
		Expect(pipe).To(Equal(pipeStruct))
	})

	It("marshals job section", func() {
		yamlBytes, err := yaml.Marshal(pipeStruct)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(yamlBytes)).To(MatchYAML(yamlStr))
	})
})
