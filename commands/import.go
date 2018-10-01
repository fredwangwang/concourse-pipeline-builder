package commands

import (
	"bytes"
	"fmt"
	"github.com/fredwangwang/concourse-pipeline-builder/builder"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Import struct {
	Config string `short:"c" long:"config" description:"path to the pipeline yaml"        required:"true" validate:"file"`
	Name   string `short:"n" long:"name"   description:"name of the imported pipeline"    required:"true" validate:"alphanum"`
	Output string `short:"o" long:"output" description:"path to the folder to be created" required:"true"`
}

var tmpl = `
package main

import (
	"fmt"
	"github.com/fredwangwang/concourse-pipeline-builder/builder"
	"gopkg.in/yaml.v2"
	"log"
)

var Pipeline = %#v

func main() {
	content, err := yaml.Marshal(Pipeline)
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

	//goPath, err := getGoPath()
	//if err != nil {
	//	return err
	//}

	//pkgPath := path.Join(goPath, i.Output)
	//fmt.Printf("going to generate pipeline in %s\n", pkgPath)

	if err := os.MkdirAll(i.Output, os.ModePerm); err != nil {
		return err
	}

	pipe := builder.Pipeline{}
	pipeBytes, err := ioutil.ReadFile(i.Config)
	if err != nil {
		return err
	}

	pipe.Name = i.Name

	err = yaml.Unmarshal(pipeBytes, &pipe)
	if err != nil {
		return err
	}

	f, err := os.Create(path.Join(i.Output, "main.go"))
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, tmpl, pipe)

	return err
}

func getGoPath() (string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("go", "env", "GOPATH")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%s %s", stderr.String(), err)
	}
	return strings.TrimSpace(stdout.String()), nil
}
