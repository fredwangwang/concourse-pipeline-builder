package builder

type Group struct {
	Name      string     `yaml:"name,omitempty"`
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

	g.Name = internal.Name
	for _, job := range internal.Jobs {
		g.Jobs = append(g.Jobs, Job{Name: job})
	}
	for _, res := range internal.Resources {
		g.Resources = append(g.Resources, Resource{Name: res})
	}

	return nil
}

func (g Group) MarshalYAML() (interface{}, error) {
	var internal _group

	internal.Name = g.Name
	for _, job := range g.Jobs {
		internal.Jobs = append(internal.Jobs, job.Name)
	}
	for _, res := range g.Resources {
		internal.Resources = append(internal.Resources, res.Name)
	}

	return internal, nil
}
