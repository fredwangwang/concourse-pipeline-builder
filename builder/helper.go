package builder

import (
	"gopkg.in/yaml.v2"
)

func yamlTransform(src interface{}, dst interface{}) error {
	content, err := yaml.Marshal(src)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, dst)
}
