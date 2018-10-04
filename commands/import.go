package commands

import (
	"fmt"
	"github.com/fredwangwang/concourse-pipeline-builder/builder"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"path"
)

type Import struct {
	Config  string `short:"c" long:"config" description:"path to the pipeline yaml"        required:"true"`
	Name    string `short:"n" long:"name"   description:"name of the imported pipeline"    required:"true"`
	Output  string `short:"o" long:"output" description:"path to the folder to be created" required:"true"`
	PerFile bool   `long:"per-file"         description:"put each type into its own file"`
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

	mainFile, err := os.Create(path.Join(i.Output, "main.go"))
	if err != nil {
		return err
	}
	defer mainFile.Close()

	pipeVarName := pipe.Generate()

	_, err = fmt.Fprint(mainFile, header)
	if err != nil {
		return err
	}

	err = i.WriteResourceTypes(mainFile)
	if err != nil {
		return err
	}

	err = i.WriteResources(mainFile)
	if err != nil {
		return err
	}

	err = i.WriteJobs(mainFile)
	if err != nil {
		return err
	}

	err = i.WriteSteps(mainFile)
	if err != nil {
		return err
	}

	err = i.WriteTasks(mainFile)
	if err != nil {
		return err
	}

	err = i.WriteGroups(mainFile)
	if err != nil {
		return err
	}

	// write Pipeline
	_, err = mainFile.WriteString(builder.GeneratedPipeline + "\n\n")
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(mainFile, fmt.Sprintf(mainFunc, pipeVarName))

	return err
}

func (i Import) WriteResourceTypes(f io.Writer) error {
	var out io.Writer

	if i.PerFile {
		file, err := os.Create(path.Join(i.Output, "resource_types.go"))
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString(`
package main

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
)
`)
		out = file
	} else {
		out = f
	}
	it := builder.ResourceTypeNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err := out.Write([]byte(kv.Value.(string) + "\n\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

func (i Import) WriteResources(f io.Writer) error {
	var out io.Writer

	if i.PerFile {
		file, err := os.Create(path.Join(i.Output, "resources.go"))
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString(`
package main

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
)
`)
		out = file
	} else {
		out = f
	}
	it := builder.ResourceNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err := out.Write([]byte(kv.Value.(string) + "\n\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

func (i Import) WriteJobs(f io.Writer) error {
	var out io.Writer

	if i.PerFile {
		file, err := os.Create(path.Join(i.Output, "jobs.go"))
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString(`
package main

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
)
`)
		out = file
	} else {
		out = f
	}
	it := builder.JobNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err := out.Write([]byte(kv.Value.(string) + "\n\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

func (i Import) WriteSteps(f io.Writer) error {
	var out io.Writer

	if i.PerFile {
		file, err := os.Create(path.Join(i.Output, "steps.go"))
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString(`
package main

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
)
`)
		out = file
	} else {
		out = f
	}
	it := builder.StepNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err := out.Write([]byte(kv.Value.(string) + "\n\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

func (i Import) WriteTasks(f io.Writer) error {
	var out io.Writer

	if i.PerFile {
		file, err := os.Create(path.Join(i.Output, "tasks.go"))
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString(`
package main

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
)
`)
		out = file
	} else {
		out = f
	}
	it := builder.TaskImageNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err := out.Write([]byte(kv.Value.(string) + "\n\n"))
		if err != nil {
			return err
		}
	}

	it = builder.TaskConfigNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err := out.Write([]byte(kv.Value.(string) + "\n\n"))
		if err != nil {
			return err
		}
	}

	return nil
}

func (i Import) WriteGroups(f io.Writer) error {
	var out io.Writer

	if i.PerFile {
		file, err := os.Create(path.Join(i.Output, "groups.go"))
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString(`
package main

import (
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
)
`)
		out = file
	} else {
		out = f
	}
	it := builder.GroupNameToBlock.IterFunc()
	for kv, ok := it(); ok; kv, ok = it() {
		_, err := out.Write([]byte(kv.Value.(string) + "\n\n"))
		if err != nil {
			return err
		}
	}

	return nil
}
