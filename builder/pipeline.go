package builder

import (
	"fmt"
	"strings"
)

type Pipeline struct {
	Name          string        `yaml:"name,omitempty"`
	ResourceTypes ResourceTypes `yaml:"resource_types,omitempty"`
	Resources     []Resource    `yaml:"resources,omitempty"`
	Jobs          []Job         `yaml:"jobs,omitempty"`
	Groups        []Group       `yaml:"groups,omitempty"`
}

// the only reason to have this is to prevent circular unmarshal
type _pipeline struct {
	Name          string        `yaml:"name,omitempty"`
	ResourceTypes ResourceTypes `yaml:"resource_types,omitempty"`
	Resources     []Resource    `yaml:"resources,omitempty"`
	Jobs          []Job         `yaml:"jobs,omitempty"`
	Groups        []Group       `yaml:"groups,omitempty"`
}

func (p *Pipeline) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var internal _pipeline

	err := unmarshal(&internal)
	if err != nil {
		return err
	}

	// linking groups
	resourcesMap := map[string]Resource{}
	jobsMap := map[string]Job{}

	for _, r := range internal.Resources {
		resourcesMap[r.Name] = r
	}
	for _, j := range internal.Jobs {
		jobsMap[j.Name] = j
	}

	for _, g := range internal.Groups {
		for i, r := range g.Resources {
			res, ok := resourcesMap[r.Name]
			if !ok {
				return fmt.Errorf("resource %s is not defined", r.Name)
			}
			g.Resources[i] = res
		}
		for i, j := range g.Jobs {
			job, ok := jobsMap[j.Name]
			if !ok {
				return fmt.Errorf("job %s is not defined", j.Name)
			}
			g.Jobs[i] = job
		}
	}

	p.Name = internal.Name
	p.ResourceTypes = internal.ResourceTypes
	p.Resources = internal.Resources
	p.Jobs = internal.Jobs
	p.Groups = internal.Groups

	return nil
}

func (p Pipeline) Generate() string {
	var parts = []string{
		"Pipeline{", // placeholder
		fmt.Sprintf("Name: \"%s\",", p.Name),
	}
	if p.ResourceTypes != nil {
		parts = append(parts, "ResourceTypes: []ResourceType{")
		for _, resType := range p.ResourceTypes {
			parts = append(parts, fmt.Sprintf("%s,", resType.Generate()))
		}
		parts = append(parts, "},")
	}
	if p.Resources != nil {
		parts = append(parts, "Resources: []Resource{")
		for _, resource := range p.Resources {
			parts = append(parts, fmt.Sprintf("%s,", resource.Generate()))
		}
		parts = append(parts, "},")
	}
	if p.Jobs != nil {
		parts = append(parts, "Jobs: []Job{")
		for _, job := range p.Jobs {
			parts = append(parts, fmt.Sprintf("%s,", job.Generate()))
		}
		parts = append(parts, "},")
	}
	if p.Groups != nil {
		parts = append(parts, "Groups: []Group{")
		for _, g := range p.Groups {
			parts = append(parts, fmt.Sprintf("%s,", g.Generate()))
		}
		parts = append(parts, "},")
	}

	// closing
	parts = append(parts, "}")

	name := fmt.Sprintf("Pipeline%s", sanitizeVarName(p.Name))
	parts[0] = fmt.Sprintf("var %s = Pipeline{", name)

	generated := strings.Join(parts, "\n")

	NameToBlock[name] = generated

	return name
}
