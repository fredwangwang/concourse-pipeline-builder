package builder

import (
	"fmt"
	"strings"
)

type Resource struct {
	Name         string                 `yaml:"name,omitempty"`
	Type         string                 `yaml:"type,omitempty"`
	Source       map[string]interface{} `yaml:"source,omitempty"`
	Version      interface{}            `yaml:"version,omitempty"`     // TODO: validate ("latest" | "every" | {version})
	CheckEvery   string                 `yaml:"check_every,omitempty"` // TODO: check is valid duration
	Tags         []string               `yaml:"tags,omitempty"`
	WebhookToken string                 `yaml:"webhook_token,omitempty"`
}

func (r Resource) Generate() string {
	var parts = []string{
		"Resource{", // placeholder
		fmt.Sprintf("Name: \"%s\",", r.Name),
		fmt.Sprintf("Type: \"%s\",", r.Type),
	}
	if r.Source != nil {
		parts = append(parts, fmt.Sprintf("Source: %#v,", r.Source))
	}
	if r.Version != nil {
		// as version can be ("every" | "latest" | obj), need to test what type is version
		if versionStr, ok := r.Version.(string); ok {
			parts = append(parts, fmt.Sprintf("Version: \"%s\",", versionStr))
		} else {
			parts = append(parts, fmt.Sprintf("Version: %#v,", r.Version))
		}
	}
	if r.CheckEvery != "" {
		parts = append(parts, fmt.Sprintf("CheckEvery: \"%s\",", r.CheckEvery))
	}
	if r.Tags != nil {
		parts = append(parts, fmt.Sprintf("Tags: %#v,", r.Tags))
	}
	if r.WebhookToken != "" {
		parts = append(parts, fmt.Sprintf("WebhookToken: \"%s\",", r.WebhookToken))
	}

	// closing
	parts = append(parts, "}")

	name := fmt.Sprintf("Resource%s", sanitizeVarName(r.Name))
	parts[0] = fmt.Sprintf("var %s = Resource{", name)

	generated := strings.Join(parts, "\n")

	NameToBlock[name] = generated

	return name
}
