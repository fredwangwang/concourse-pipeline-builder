package main

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"

	"fmt"
	"gopkg.in/yaml.v2"
	"log"
)

var ResourceTypePivnet = ResourceType{
	Name: "pivnet",
	Type: "docker-image",
	Source: map[string]interface{}{
		"repository": "pivotalcf/pivnet-resource",
		"tag":        "latest-final",
	},
}

var ResourcePcfPipelines = Resource{
	Name: "pcf-pipelines",
	Type: "git",
	Source: map[string]interface{}{
		"uri":    "git@github.com:pivotal-cf/pcf-pipelines.git",
		"branch": "master",
	},
}

var ResourcePASTile = Resource{
	Name: "tile",
	Type: "pivnet",
	Source: map[string]interface{}{
		"api_token":    "token",
		"product_slug": "elastic-runtime",
	},
}

var ResourceSchedule = Resource{
	Name: "schedule",
	Type: "time",
	Source: map[string]interface{}{
		"interval": "30m",
		"start":    "12:00 AM",
		"stop":     "11:59 PM",
		"location": "America/Los_Angeles",
		"days":     []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
	},
}

func main() {
	a := Pipeline{
		//Name: "testing123",
		ResourceTypes: []ResourceType{
			ResourceTypePivnet,
		},
		Resources: []Resource{
			ResourcePcfPipelines,
			ResourcePASTile,
			ResourceSchedule,
		},
		Jobs: []Job{
			{
				Name: "regulator",
				Plan: []Step{
					StepGet{
						Get:     "schedule",
						Trigger: true,
						StepHook: StepHook{
							Tags: []string{
								"test",
							},
						},
					},
					StepGet{
						Get: "tile",
						Params: map[string]interface{}{
							"globs": []string{},
						},
					},
				},
			},
		},
		Groups: nil,
	}

	content, err := yaml.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
