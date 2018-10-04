package builder_test

import (
	. "github.com/onsi/ginkgo"
	"gopkg.in/yaml.v2"

	. "github.com/onsi/gomega"

	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
)

var _ = Describe("Pipeline", func() {
	var yamlStr = `
name: a-pipeline
resource_types:
- name: pivnet
  type: docker-image
  source:
    repository: pivotalcf/pivnet-resource
    tag: latest-final
- name: awesome-resource
  type: docker-image
  source:
    repository: awesome/resource
  privileged: true
  check_every: 5m
  tags: [ "awesome-worker" ]
resources:
- name: some-pivnet-product
  type: pivnet
  source:
    api_token: token
- name: awesome-product
  type: awesome-resource
  source:
    awesome: true
  version: latest
  webhook_token: localhost/check_if_awesome
jobs:
- name: job1
  plan:
  - aggregate:
    - get: res1
      resource: some-pivnet-product
    - get: res2
      resource: awesome-product
      timeout: 24h
      attempts: 12345
  - task: task1
    file: some/file.yml
    on_success:
      get: some-pivnet-product
    on_failure:
      put: awesome-product
      params:
        awesome: false
  - put: awesome-product
- name: job2
  plan:
  - get: awesome-product
groups:
- name: all
  jobs:
  - job1
  - job2
  resources:
  - some-pivnet-product
  - awesome-product
- name: j2
  jobs:
  - job2
  resources:
  - awesome-product
`

	var resourceTypePivnet = ResourceType{
		Name: "pivnet",
		Type: "docker-image",
		Source: map[string]interface{}{
			"repository": "pivotalcf/pivnet-resource",
			"tag":        "latest-final",
		},
	}

	var resourceTypeAwesome = ResourceType{
		Name: "awesome-resource",
		Type: "docker-image",
		Source: map[string]interface{}{
			"repository": "awesome/resource",
		},
		Privileged: true,
		CheckEvery: "5m",
		Tags:       []string{"awesome-worker"},
	}

	var resourcePivnetProduct = Resource{
		Name: "some-pivnet-product",
		Type: "pivnet",
		Source: map[string]interface{}{
			"api_token": "token",
		},
	}

	var resourceAwesomeProduct = Resource{
		Name: "awesome-product",
		Type: "awesome-resource",
		Source: map[string]interface{}{
			"awesome": true,
		},
		Version:      "latest",
		WebhookToken: "localhost/check_if_awesome",
	}

	var job1 = Job{
		Name: "job1",
		Plan: []Step{
			StepAggregate{
				Aggregate: Steps{
					StepGet{
						Get:      "res1",
						Resource: "some-pivnet-product",
					},
					StepGet{
						Get:      "res2",
						Resource: "awesome-product",
						StepHook: StepHook{
							Timeout:  "24h",
							Attempts: 12345,
						},
					},
				},
			},
			StepTask{
				Task: "task1",
				File: "some/file.yml",
				StepHook: StepHook{
					OnSuccess: StepGet{
						Get: "some-pivnet-product",
					},
					OnFailure: StepPut{
						Put: "awesome-product",
						Params: map[string]interface{}{
							"awesome": false,
						},
					},
				},
			},
			StepPut{
				Put: "awesome-product",
			},
		},
	}

	var job2 = Job{
		Name: "job2",
		Plan: Steps{
			StepGet{
				Get: "awesome-product",
			},
		},
	}

	var groupAll = Group{
		Name:      "all",
		Jobs:      []Job{job1, job2},
		Resources: []Resource{resourcePivnetProduct, resourceAwesomeProduct},
	}

	var groupJ2 = Group{
		Name:      "j2",
		Jobs:      []Job{job2},
		Resources: []Resource{resourceAwesomeProduct},
	}

	var pipeStruct = Pipeline{
		Name:          "a-pipeline",
		ResourceTypes: []ResourceType{resourceTypePivnet, resourceTypeAwesome},
		Resources:     []Resource{resourcePivnetProduct, resourceAwesomeProduct},
		Jobs:          []Job{job1, job2},
		Groups:        []Group{groupAll, groupJ2},
	}

	It("unmarshals the pipeline section", func() {
		var pipe Pipeline
		err := yaml.Unmarshal([]byte(yamlStr), &pipe)
		Expect(err).NotTo(HaveOccurred())
		Expect(pipe).To(Equal(pipeStruct))
	})

	It("marshals pipeline section", func() {
		yamlBytes, err := yaml.Marshal(pipeStruct)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(yamlBytes)).To(MatchYAML(yamlStr))
	})

	It("generates proper code section", func() {
		step1 := Pipeline{
			Name:          "p-1",
			ResourceTypes: []ResourceType{{Name: "rt1", Type: "docker"}},
			Resources:     []Resource{{Name: "r1", Type: "rt1"}},
			Jobs:          []Job{{Name: "j1"}},
			Groups:        []Group{{Name: "g1", Jobs: []Job{{Name: "j1"}}}},
		}

		expected := `var Pipelinep1 = Pipeline{
Name: "p-1",
ResourceTypes: []ResourceType{
ResourceTypert1,
},
Resources: []Resource{
Resourcer1,
},
Jobs: []Job{
Jobj1,
},
Groups: []Group{
Groupg1,
},
}`

		stepName := step1.Generate()
		Expect(stepName).To(Equal("Pipelinep1"))
		result := GeneratedPipeline
		GinkgoWriter.Write([]byte(result))
		Expect(result).To(Equal(expected))
	})
})
