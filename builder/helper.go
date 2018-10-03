package builder

import (
	"gopkg.in/yaml.v2"
	"regexp"
)

func yamlTransform(src interface{}, dst interface{}) error {
	content, err := yaml.Marshal(src)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, dst)
}

var reg = regexp.MustCompile("[^a-zA-Z0-9]+")

func sanitizeVarName(s string) string {
	return reg.ReplaceAllString(s, "")
}
