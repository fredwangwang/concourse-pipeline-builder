package builder

type Resource struct {
	Name         string                 `yaml:"name,omitempty"`
	Type         string                 `yaml:"type,omitempty"`
	Source       map[string]interface{} `yaml:"source,omitempty"`
	Version      interface{}            `yaml:"version,omitempty"` // TODO: maybe map[string]interface{} ?
	CheckEvery   string                 `yaml:"check_every,omitempty"`
	Tags         []string               `yaml:"tags,omitempty"`
	WebhookToken string                 `yaml:"webhook_token,omitempty"`
}
