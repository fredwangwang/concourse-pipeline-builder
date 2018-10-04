package commands

import (
	"fmt"
	"github.com/fredwangwang/concourse-pipeline-builder/builder"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
)

type Import struct {
	Config string `short:"c" long:"config" description:"path to the pipeline yaml"        required:"true"`
	Name   string `short:"n" long:"name"   description:"name of the imported pipeline"    required:"true"`
	Output string `short:"o" long:"output" description:"path to the folder to be created" required:"true"`
}

const header = `
package main

import (
	"fmt"
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	"gopkg.in/yaml.v2"
	"log"
)
`

const mainFunc = `
func main() {
	content, err := yaml.Marshal(%s)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(string(content))
}
`

func (i *Import) Execute(args []string) error {
	if err := validator.New().Struct(i); err != nil {
		return err
	}

	if err := os.MkdirAll(i.Output, os.ModePerm); err != nil {
		return err
	}

	pipe := builder.Pipeline{}
	pipeBytes, err := ioutil.ReadFile(i.Config)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(pipeBytes, &pipe)
	if err != nil {
		return err
	}

	pipe.Name = i.Name

	f, err := os.Create(path.Join(i.Output, "main.go"))
	if err != nil {
		return err
	}
	defer f.Close()

	pipeVarName := pipe.Generate()

	_, err = fmt.Fprint(f, header)
	if err != nil {
		return err
	}

	for _, block := range builder.NameToBlock {
		_, err = f.WriteString(block + "\n\n")
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(f, fmt.Sprintf(mainFunc, pipeVarName))

	return err
}
