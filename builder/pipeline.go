package builder

type Pipeline struct {
	Name          string         `yaml:"name,omitempty"`
	ResourceTypes []ResourceType `yaml:"resource_types,omitempty"`
	Resources     []Resource     `yaml:"resources,omitempty"`
	Jobs          []Job          `yaml:"jobs,omitempty"`
	Groups        []Group        `yaml:"groups,omitempty"`
}
