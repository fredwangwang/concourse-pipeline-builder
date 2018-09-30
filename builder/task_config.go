package builder

// TODO: add validate

type TaskConfig struct {
	Platform      string                 `yaml:"platform,omitempty"`
	ImageResource TaskImageResource      `yaml:"image_resource,omitempty"`
	RootfsUri     string                 `yaml:"rootfs_uri,omitempty"`
	Inputs        []TaskInput            `yaml:"inputs,omitempty"`
	Outputs       []TaskOutput           `yaml:"outputs,omitempty"`
	Caches        []TaskCache            `yaml:"caches,omitempty"`
	Run           TaskRun                `yaml:"run,omitempty"`
	Params        map[string]interface{} `yaml:"params,omitempty"`
}

type TaskImageResource struct {
	Type    string                 `yaml:"type,omitempty"`
	Source  map[string]interface{} `yaml:"source,omitempty"`
	Params  map[string]interface{} `yaml:"params,omitempty"`
	Version interface{}            `yaml:"version,omitempty"`
}

type TaskInput struct {
	Name     string `yaml:"name,omitempty"`
	Path     string `yaml:"path,omitempty"`
	Optional bool   `yaml:"optional,omitempty"`
}

type TaskOutput struct {
	Name string `yaml:"name,omitempty"`
	Path string `yaml:"path,omitempty"`
}

type TaskCache struct {
	Path string `yaml:"path,omitempty"`
}

type TaskRun struct {
	Path string   `yaml:"path,omitempty"`
	Args []string `yaml:"args,omitempty"`
	Dir  string   `yaml:"dir,omitempty"`
	User string   `yaml:"user,omitempty"`
}
