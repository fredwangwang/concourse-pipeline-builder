package main

import (
	"fmt"
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
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
	stepSchedule := StepGet{
		Get:     "schedule",
		Trigger: true,
		StepHook: StepHook{
			Tags: []string{
				"test",
			},
		},
	}

	stepGetTile := StepGet{
		Get: "tile",
		Params: map[string]interface{}{
			"globs": []string{},
		},
	}

	job1 := Job{
		Name: "regulator",
		Plan: []Step{stepSchedule, stepGetTile},
	}

	a := Pipeline{
		Name: "",
		ResourceTypes: []ResourceType{
			ResourceTypePivnet,
		},
		Resources: []Resource{
			ResourcePASTile,
			ResourceSchedule,
		},
		Jobs: []Job{
			job1,
		},
		Groups: []Group{
			{
				Name: "a-group",
				Jobs: []Job{
					job1,
				},
			},
		},
	}

	content, err := yaml.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
