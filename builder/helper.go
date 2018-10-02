package builder

import (
	"gopkg.in/yaml.v2"
	"hash/fnv"
)

func yamlTransform(src interface{}, dst interface{}) error {
	content, err := yaml.Marshal(src)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(content, dst)
}

func hashBytes(b []byte) uint32 {
	h := fnv.New32a()
	h.Write(b)
	return h.Sum32()
}

func hashString(s string) uint32 {
	return hashBytes([]byte(s))
}
