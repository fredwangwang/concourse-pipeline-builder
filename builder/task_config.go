package builder

import (
	"fmt"
	"github.com/mitchellh/hashstructure"
	"log"
	"strings"
)

// TODO: add validate

type TaskConfig struct {
	Platform      string                 `yaml:"platform,omitempty"`
	ImageResource *TaskImageResource     `yaml:"image_resource,omitempty"`
	RootfsUri     string                 `yaml:"rootfs_uri,omitempty"`
	Inputs        []TaskInput            `yaml:"inputs,omitempty"`
	Outputs       []TaskOutput           `yaml:"outputs,omitempty"`
	Caches        []TaskCache            `yaml:"caches,omitempty"`
	Run           *TaskRun               `yaml:"run,omitempty"`
	Params        map[string]interface{} `yaml:"params,omitempty"`
}

type TaskImageResource struct {
	Type    string                 `yaml:"type,omitempty"`
	Source  map[string]interface{} `yaml:"source,omitempty"`
	Params  map[string]interface{} `yaml:"params,omitempty"`
	Version map[string]interface{} `yaml:"version,omitempty"`
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

func (t TaskConfig) Generate() string {
	var parts = []string{
		"TaskConfig:{", // placeholder
		fmt.Sprintf("Platform: \"%s\",", t.Platform),
	}
	if t.ImageResource != nil {
		parts = append(parts, fmt.Sprintf("ImageResource: &%s,", t.ImageResource.Generate()))
	}
	if t.RootfsUri != "" {
		parts = append(parts, fmt.Sprintf("RootfsUri: \"%s\",", t.RootfsUri))
	}
	if t.Inputs != nil {
		parts = append(parts, fmt.Sprintf("Inputs: []TaskInput{"))
		for _, input := range t.Inputs {
			line := fmt.Sprintf("{Name: \"%s\", Path: \"%s\", Optional: %v},",
				input.Name, input.Path, input.Optional)
			parts = append(parts, line)
		}
		parts = append(parts, "},")
	}
	if t.Outputs != nil {
		parts = append(parts, fmt.Sprintf("Outputs: []TaskOutput{"))
		for _, output := range t.Outputs {
			line := fmt.Sprintf("{Name: \"%s\", Path: \"%s\"},",
				output.Name, output.Path)
			parts = append(parts, line)
		}
		parts = append(parts, "},")
	}
	if t.Caches != nil {
		parts = append(parts, fmt.Sprintf("Caches: []TaskCache{"))
		for _, cache := range t.Caches {
			line := fmt.Sprintf("{Path: \"%s\"},", cache.Path)
			parts = append(parts, line)
		}
		parts = append(parts, "},")
	}
	if t.Run != nil {
		parts = append(parts, fmt.Sprintf("Run: &TaskRun{"))
		parts = append(parts, fmt.Sprintf("Path: \"%s\",", t.Run.Path))
		parts = append(parts, fmt.Sprintf("Dir: \"%s\",", t.Run.Dir))
		parts = append(parts, fmt.Sprintf("User: \"%s\",", t.Run.User))
		parts = append(parts, fmt.Sprintf("Args: %#v,", t.Run.Args))
		parts = append(parts, "},")
	}
	if t.Params != nil {
		parts = append(parts, fmt.Sprintf("Params: %#v,", t.Params))
	}

	// closing
	parts = append(parts, "}")

	// get name
	// hash is deterministic. Given the same struct, the hash is always the same.
	hash, err := hashstructure.Hash(t, nil)
	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("TaskConfig%x", hash)
	parts[0] = fmt.Sprintf("var %s = TaskConfig{", name)

	generated := strings.Join(parts, "\n")

	NameToBlock[name] = generated

	return name
}

func (t TaskImageResource) Generate() string {
	var parts = []string{
		"TaskImageResource:{", // placeholder
	}
	parts = append(parts, fmt.Sprintf("Type: \"%s\",", t.Type))
	parts = append(parts, fmt.Sprintf("Source: %#v,", t.Source))
	if t.Params != nil {
		parts = append(parts, fmt.Sprintf("Params: %#v,", t.Params))
	}
	if t.Version != nil {
		parts = append(parts, fmt.Sprintf("Version: %#v,", t.Version))
	}
	parts = append(parts, "}")

	hash, err := hashstructure.Hash(t, nil)
	if err != nil {
		log.Fatal(err)
	}

	name := fmt.Sprintf("TaskImageResource%x", hash)
	parts[0] = fmt.Sprintf("var %s = TaskImageResource{", name)

	generated := strings.Join(parts, "\n")

	NameToBlock[name] = generated

	return name
}
