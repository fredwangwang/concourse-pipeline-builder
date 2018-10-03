package builder_test

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Describe("ResourceType", func() {
	var yamlStr = `
resource_types:
- name: pivnet
  type: something
  source:
    repository: pivotalcf/pivnet-resource
    tag: latest-final
`

	var pipeStruct = Pipeline{
		ResourceTypes: ResourceTypes{
			ResourceType{
				Name: "pivnet",
				Type: "something",
				Source: map[string]interface{}{
					"repository": "pivotalcf/pivnet-resource",
					"tag":        "latest-final",
				},
			},
		},
	}

	It("unmarshals the resource types section", func() {
		var pipe Pipeline
		err := yaml.Unmarshal([]byte(yamlStr), &pipe)
		Expect(err).NotTo(HaveOccurred())
		Expect(pipe).To(Equal(pipeStruct))
	})

	It("marshals resource types section", func() {
		yamlBytes, err := yaml.Marshal(pipeStruct)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(yamlBytes)).To(MatchYAML(yamlStr))
	})

	It("generates proper code section", func() {
		step1 := ResourceType{
			Name: "res1",
			Type: "some",
			Source: map[string]interface{}{
				"a": "b",
			},
			Privileged: true,
			Params: map[string]interface{}{
				"b": "c",
			},
			CheckEvery: "4m",
			Tags:       []string{"a", "z"},
		}

		expected := `var ResourceTyperes1 = ResourceType{
Name: "res1",
Type: "some",
Source: map[string]interface {}{"a":"b"},
Privileged: true,
Params: map[string]interface {}{"b":"c"},
CheckEvery: "4m",
Tags: []string{"a", "z"},
}`

		stepName := step1.Generate()
		result, ok := NameToBlock[stepName]
		Expect(ok).To(BeTrue())
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(Equal(expected))
	})
})
