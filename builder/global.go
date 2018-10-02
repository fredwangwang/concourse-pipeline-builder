package builder

import "math/rand"

type Generator interface {
	Generate() (interface{}, string)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var IndexBlock = map[uint32]string{}
var VarNameIndex = map[string]interface{}{}
var StepNameToBlock = map[string]string{}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
