package builder

import (
	"gopkg.in/go-playground/validator.v9"
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
