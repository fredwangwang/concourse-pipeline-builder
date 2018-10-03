package builder_test

import (
	. "github.com/onsi/ginkgo"
	"gopkg.in/yaml.v2"

	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/gomega"
)

var _ = Describe("Group", func() {
	var yamlStr = `
jobs:
- name: j1
groups:
- name: a-group
  jobs: [ "j1" ]
`

	var pipeStruct = Pipeline{
		Jobs: []Job{
			{
				Name: "j1",
			},
		},
		Groups: []Group{
			{
				Name: "a-group",
				Jobs: []Job{
					{Name: "j1"},
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

	It("validates", func() {
		group := Group{}
		str := "{}"

		err := yaml.Unmarshal([]byte(str), &group)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("'Name' failed on the 'required' tag"))

		_, err = yaml.Marshal(group)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("'Name' failed on the 'required' tag"))
	})

	It("generates proper code section", func() {
		step1 := Group{
			Name:      "g1",
			Jobs:      []Job{{Name: "j1"}},
			Resources: []Resource{{Name: "r1"}},
		}

		expected := `var Groupg1 = Group{
Name: "g1",
Jobs: []Job{
Jobj1,
},
Resources: []Resource{
Resourcer1,
},
}`

		stepName := step1.Generate()
		result, ok := NameToBlock[stepName]
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(Equal(expected))
	})
})
