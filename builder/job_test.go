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

	It("generates proper code section", func() {
		step1 := Job{
			Name: "jb1",
			Plan: Steps{
				StepGet{Get: "some"},
				StepPut{Put: "else"},
			},
			Serial:               true,
			SerialGroups:         []string{"abc"},
			BuildLogsToRetain:    123,
			MaxInFlight:          321,
			Public:               true,
			DisableManualTrigger: true,
			Interruptible:        true,
			OnSuccess:            StepGet{Get: "OnSuccess"},
			OnFailure:            StepGet{Get: "OnFailure"},
			OnAbort:              StepGet{Get: "OnAbort"},
			Ensure:               StepGet{Get: "Ensure"},
		}

		expected := `var Jobjb1 = Job{
Name: "jb1",
Plan: Steps{
StepGetsome8d05d7738e82dae,
StepPutelsea6bc5d26f886948,
},
Serial: true,
SerialGroups: []string{"abc"},
BuildLogsToRetain: 123,
MaxInFlight: 321,
Public: true,
DisableManualTrigger: true,
Interruptible: true,
OnSuccess: StepGetOnSuccessd4ce2d2696ef7d12,
OnFailure: StepGetOnFailure41ffd933a5738f0d,
OnAbort: StepGetOnAbort4e9644bd74da6ad2,
Ensure: StepGetEnsure693c228a58fc8a26,
}`

		stepName := step1.Generate()
		result, ok := JobNameToBlock.Get(stepName)
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result.(string)))
		Expect(result).To(Equal(expected))
	})
})
