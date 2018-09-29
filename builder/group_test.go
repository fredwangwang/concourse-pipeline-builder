package builder_test

import (
	. "github.com/onsi/ginkgo"
	"gopkg.in/yaml.v2"

	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/gomega"
)

var _ = Describe("Group", func() {
	var yamlStr = `
groups:
- name: a-group
  jobs: [ "j1", "j2" ]
`

	var pipeStruct = Pipeline{
		Groups: []Group{
			{
				Name: "a-group",
				Jobs: []Job{
					{Name: "j1"}, {Name: "j2"},
				},
			},
		},
	}

	It("unmarshals groups section", func() {
		var pipe Pipeline
		err := yaml.Unmarshal([]byte(yamlStr), &pipe)
		Expect(err).NotTo(HaveOccurred())
		Expect(pipe).To(Equal(pipeStruct))
	})

	It("marshals groups section", func() {
		yamlBytes, err := yaml.Marshal(pipeStruct)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(yamlBytes)).To(MatchYAML(yamlStr))
	})
})
