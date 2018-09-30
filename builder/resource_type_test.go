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

	It("validates", func() {
		rt := ResourceType{}
		err := rt.Validate()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("'Name' failed on the 'required' tag"))
		Expect(err.Error()).To(ContainSubstring("'Type' failed on the 'required' tag"))
	})
})
