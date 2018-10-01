package builder

type Generator interface {
	Generate() (interface{}, string)
}

var IndexBlock = map[interface{}]string{}
var VarNameIndex = map[string]interface{}{}
