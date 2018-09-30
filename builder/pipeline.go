package builder

import "fmt"

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
