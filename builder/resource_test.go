package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("Resource", func() {
	var yamlStr = `
resources:
- name: test-resource
  type: git
  source:
    uri: git@github.com/fredwangwang/orderedmap
    branch: master
  check_every: 1m
  webhook_token: localhost/check_test-resource
`

	var pipeStruct = Pipeline{
		Resources: []Resource{
			{
				Name: "test-resource",
				Type: "git",
				Source: map[string]interface{}{
					"uri":    "git@github.com/fredwangwang/orderedmap",
					"branch": "master",
				},
				CheckEvery:   "1m",
				WebhookToken: "localhost/check_test-resource",
			},
		},
	}

	It("unmarshals the resource section", func() {
		var pipe Pipeline
		err := yaml.Unmarshal([]byte(yamlStr), &pipe)
		Expect(err).NotTo(HaveOccurred())
		Expect(pipe).To(Equal(pipeStruct))
	})

	It("marshals resource section", func() {
		yamlBytes, err := yaml.Marshal(pipeStruct)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(yamlBytes)).To(MatchYAML(yamlStr))
	})
})
