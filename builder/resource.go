package builder

type Resource struct {
	Name         string                 `yaml:"name,omitempty"`
	Type         string                 `yaml:"type,omitempty"`
	Source       map[string]interface{} `yaml:"source,omitempty"`
	Version      interface{}            `yaml:"version,omitempty"` // TODO: validate ("latest" | "every" | {version})
	CheckEvery   string                 `yaml:"check_every,omitempty"` // TODO: check is valid duration
	Tags         []string               `yaml:"tags,omitempty"`
	WebhookToken string                 `yaml:"webhook_token,omitempty"`
}