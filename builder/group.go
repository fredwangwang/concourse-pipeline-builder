package builder

type Group struct {
	Name      string   `yaml:"name,omitempty"`
	Jobs      []string `yaml:"jobs,omitempty"`      //[]Job		// TODO: see if it is possible to use []Job instead of []string for schema validation
	Resources []string `yaml:"resources,omitempty"` //[]Resource
}
