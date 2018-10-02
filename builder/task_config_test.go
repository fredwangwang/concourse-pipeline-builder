package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/gomega"

	. "github.com/onsi/ginkgo"
	"gopkg.in/yaml.v2"
)

var _ = Describe("TaskConfig", func() {
	var yamlStr = `
jobs:
- name: hw
  plan:
  - task: hello
    config:
      platform: linux
      image_resource:
        type: docker-image
        source:
          repository: ubuntu
      inputs:
      - name: config
      outputs:
      - name: config-updated
      caches:
      - path: temp-res
      run:
        path: bash
        args:
        - -c
        - |

          set -eux
          echo "$HELLO_STR"
      params:
        HELLO_STR:
    params:
      HELLO_STR: hello-world
    attempts: 2
`

	var pipeStruct = Pipeline{
		Jobs: []Job{
			{
				Name: "hw",
				Plan: Steps{
					StepTask{
						Task: "hello",
						Config: &TaskConfig{
							Platform: "linux",
							ImageResource: &TaskImageResource{
								Type: "docker-image",
								Source: map[string]interface{}{
									"repository": "ubuntu",
								},
							},
							Inputs: []TaskInput{
								{
									Name: "config",
								},
							},
							Outputs: []TaskOutput{
								{
									Name: "config-updated",
								},
							},
							Caches: []TaskCache{
								{
									Path: "temp-res",
								},
							},
							Run: &TaskRun{
								Path: "bash",
								Args: []string{
									`-c`,
									`
set -eux
echo "$HELLO_STR"
`,
								},
							},
							Params: map[string]interface{}{
								"HELLO_STR": nil,
							},
						},
						Params: map[string]interface{}{
							"HELLO_STR": "hello-world",
						},
						StepHook: StepHook{
							Attempts: 2,
						},
					},
				},
			},
		},
	}

	It("unmarshals the task config section", func() {
		var pipe Pipeline
		err := yaml.Unmarshal([]byte(yamlStr), &pipe)
		Expect(err).NotTo(HaveOccurred())
		Expect(pipe).To(Equal(pipeStruct))
	})

	It("marshals the task config section", func() {
		yamlBytes, err := yaml.Marshal(pipeStruct)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(yamlBytes)).To(MatchYAML(yamlStr))
	})

	It("generates proper code section", func() {
		imageResource := TaskImageResource{
			Type: "docker",
			Source: map[string]interface{}{
				"from": "something",
			},
			Params: map[string]interface{}{
				"to": "somewhere",
			},
			Version: map[string]interface{}{
				"i":    "actually",
				"dont": "know",
				"if":   "this",
				"is":   "been used",
			},
		}

		imageName := imageResource.Generate()
		result, ok := StepNameToBlock[imageName]
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(ContainSubstring("var TaskImageResourceaf6e4556101cbde"))

		taskConfig := TaskConfig{
			Platform:      "linux",
			ImageResource: &imageResource,
			RootfsUri:     "aaa",
			Inputs: []TaskInput{
				{
					Name:     "in",
					Path:     "path",
					Optional: true,
				},
			},
			Outputs: []TaskOutput{
				{
					Name: "out",
					Path: "path1",
				},
			},
			Caches: []TaskCache{
				{
					Path: "cpath",
				},
			},
			Run: &TaskRun{
				Path: "bash",
				Args: []string{
					"bcd",
				},
				Dir:  "",
				User: "",
			},
			Params: map[string]interface{}{
				"p1": "val",
			},
		}

		expectedTaskConfig := `var TaskConfig3721a40786706e0a = TaskConfig{
Platform: "linux",
ImageResource: &TaskImageResourceaf6e4556101cbde,
RootfsUri: "aaa",
Inputs: []TaskInput{
{Name: "in", Path: "path", Optional: true},
},
Outputs: []TaskOutput{
{Name: "out", Path: "path1"},
},
Caches: []TaskCache{
{Path: "cpath"},
},
Run: &TaskRun{
Path: "bash",
Dir: "",
User: "",
Args: []string{"bcd"},
},
Params: map[string]interface {}{"p1":"val"},
}`

		stepName := taskConfig.Generate()
		result, ok = StepNameToBlock[stepName]
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(Equal(expectedTaskConfig))
	})
})
