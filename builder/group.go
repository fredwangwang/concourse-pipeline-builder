package builder

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type Group struct {
	Name      string     `yaml:"name,omitempty" validate:"required"`
	Jobs      []Job      `yaml:"jobs,omitempty"`
	Resources []Resource `yaml:"resources,omitempty"`
}

type _group struct {
	Name      string   `yaml:"name,omitempty"`
	Jobs      []string `yaml:"jobs,omitempty"`
	Resources []string `yaml:"resources,omitempty"`
}

func (g *Group) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var internal _group

	err := unmarshal(&internal)
	if err != nil {
		return err
	}

	// the following steps convert []string to []job / []resource, with only name set.
	// it is not particularly useful at this point, but it will be referenced when
	// unmarshaling the pipeline, which will validate and link to the resources defined.
	// there is no need to validate jobs & resources at this level, defer that to pipeline.
	g.Name = internal.Name
	for _, job := range internal.Jobs {
		g.Jobs = append(g.Jobs, Job{Name: job})
	}
	for _, res := range internal.Resources {
		g.Resources = append(g.Resources, Resource{Name: res})
	}

	validate := validator.New()
	return validate.Struct(g)
}

func (g Group) MarshalYAML() (interface{}, error) {
	var internal _group

	validate := validator.New()
	err := validate.Struct(g)
	if err != nil {
		return nil, err
	}

	// convert []job / []resource to []string for yaml generation.
	internal.Name = g.Name
	for _, job := range g.Jobs {
		internal.Jobs = append(internal.Jobs, job.Name)
	}
	for _, res := range g.Resources {
		internal.Resources = append(internal.Resources, res.Name)
	}

	return internal, nil
}

func (g Group) Generate() string {
	var parts = []string{
		"Group{", // placeholder
		fmt.Sprintf("Name: \"%s\",", g.Name),
	}

	// the generation of groups depends on the parsed jobs and resources. As it makes
	// no sense that a job is defined only in group section not in jobs section as well.
	// in order to correctly linking the jobs and resources, this method must be called
	// after the generation of jobs and resources. Otherwise exception will be thrown.
	// eh actually not, I realized as I typed out these things... hash solves the problem,
	// thats why I used hash in the first place... It will have collision for the two same
	// code block, resulting only one instance in the global map.
	if g.Jobs != nil {
		parts = append(parts, "Jobs: []Job{")
		for _, job := range g.Jobs {
			parts = append(parts, fmt.Sprintf("%s,", job.Generate()))
		}
		parts = append(parts, "},")
	}
	if g.Resources != nil {
		parts = append(parts, "Resources: []Resource{")
		for _, resource := range g.Resources {
			parts = append(parts, fmt.Sprintf("%s,", resource.Generate()))
		}
		parts = append(parts, "},")
	}

	// closing
	parts = append(parts, "}")

	name := fmt.Sprintf("Group%s", sanitizeVarName(g.Name))
	parts[0] = fmt.Sprintf("var %s = Group{", name)

	generated := strings.Join(parts, "\n")

	NameToBlock[name] = generated

	return name
}
