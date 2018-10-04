package commands

import (
	"fmt"
	"github.com/fredwangwang/concourse-pipeline-builder/builder"
	"github.com/fredwangwang/orderedmap"
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

	var it func() (*orderedmap.KVPair, bool)

	// write ResourceTypes
	it = builder.ResourceTypeNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err = f.WriteString(kv.Value.(string) + "\n\n")
		if err != nil {
			return err
		}
	}

	// write Resources
	it = builder.ResourceNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err = f.WriteString(kv.Value.(string) + "\n\n")
		if err != nil {
			return err
		}
	}

	// write Jobs
	it = builder.JobNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err = f.WriteString(kv.Value.(string) + "\n\n")
		if err != nil {
			return err
		}
	}

	// write Steps
	it = builder.StepNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err = f.WriteString(kv.Value.(string) + "\n\n")
		if err != nil {
			return err
		}
	}

	// write TaskConfigs
	it = builder.TaskConfigNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err = f.WriteString(kv.Value.(string) + "\n\n")
		if err != nil {
			return err
		}
	}

	// write TaskImages
	it = builder.TaskImageNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err = f.WriteString(kv.Value.(string) + "\n\n")
		if err != nil {
			return err
		}
	}

	// write Groups
	it = builder.GroupNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err = f.WriteString(kv.Value.(string) + "\n\n")
		if err != nil {
			return err
		}
	}

	// write Pipeline
	_, err = f.WriteString(builder.GeneratedPipeline + "\n\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(f, fmt.Sprintf(mainFunc, pipeVarName))

	return err
}
