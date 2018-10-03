# Concourse Pipeline Builder

Pipeline as code, allows you to write pipeline in golang, which is a real programming language
you can compose rather than patching together a thousand line yaml.

## Install
```bash
$ go get github.com/fredwangwang/concourse-pipeline-builder
$ concourse-pipeline-builder -h
2018/10/02 22:53:55 Usage:
  concourse-pipeline-builder [OPTIONS] <import>

Help Options:
  -h, --help  Show this help message

Available commands:
  import  import the existing pipeline
```

## Example Usage:
1. Import the existing pipeline  
Imagine you have an existing pipeline and would like to try this tool out, you can easily
import your pipeline. Simply run:  
`concourse-pipeline-builder import -c path/to/pipeline.yml -o output/dir -n name-of-pipeline`,
you would get a pipeline in Go! The generated pipeline would be in `output/dir/main.go`.  
 
1. Starting from scratch  
It's easy. All you need is the following boilerplate code and it will get you started.  
```go
package main

import (
	"fmt"
	. "github.com/fredwangwang/concourse-pipeline-builder/builder"
	"gopkg.in/yaml.v2"
	"log"
)

var pipe = Pipeline{
	// name key is not picked up by concourse
	// it is reserved for the code generation purpose
	Name: "Sample-pipeline",
	ResourceTypes: []ResourceType{},
	Resources:[]Resource{},
	Jobs:[]Job{},
	Groups:[]Group{},
}

func main() {
	content, err := yaml.Marshal(pipe)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
```

### Note:
To get the yaml equivalent of the _codified_ pipeline, do `go run path/to/main.go`.
This allows you to set the pipeline by running:
`fly -t target sp -p pipeline -c <(go run path/to/main.go)`


## Why
Concourse CI is a powerful tool, but building the pipeline may not be a pleasant process.
Especially for large pipelines there might be thousand lines of yaml, it is simply not manageable.

There are various ways to make the process better, including yaml templating, bosh ops file, etc.
They all works at some degree, but IMO, they do not handle a root problem: using yaml as a DSL.

Yaml is great for human-readable data serialization solution. It is much easier to interact with
compared with json, but does not mean it is great for being a carrier for structured language.

Some pros & cons I could think of using yaml as a carrier for DSL in general:

#### Pro:
* Widely used
* Structurally simple
* Easy to change quickly

#### Con:
* Schema-less
* Not very easy to lint
* Not easy to compose
* Hard to maintain

So I want to create a tool (basically for fun) to allow compiling a pipeline
using a more structured language (golang), which allows:
* Schema enforcing
* Syntax highlighting
* Composability
* Reusability

## State
- [x] Basic structures
- [ ] Validation when marshaling 
- [ ] Make it a cli (Partially done)
- [ ] Better structure the generated code
- [ ] todo?

There are lots of improving opportunities to this command, right now it is a working prototype 
of how I want it to work.

## Example
See `example/pipeline.go`, it will generate the following pipeline:
```yaml
resource_types:
- name: pivnet
  type: docker-image
  source:
    repository: pivotalcf/pivnet-resource
    tag: latest-final
resources:
- name: tile
  type: pivnet
  source:
    api_token: token
    product_slug: elastic-runtime
- name: schedule
  type: time
  source:
    days:
    - Sunday
    - Monday
    - Tuesday
    - Wednesday
    - Thursday
    - Friday
    - Saturday
    interval: 30m
    location: America/Los_Angeles
    start: 12:00 AM
    stop: 11:59 PM
jobs:
- name: regulator
  plan:
  - get: schedule
    trigger: true
    tags:
    - test
  - get: tile
    params:
      globs: []
  - task: hello
    config:
      platform: linux
      image_resource:
        type: docker-image
        source:
          repository: ubuntu
      inputs:
      - name: config
      outputs:
      - name: config-updated
      caches:
      - path: temp-res
      run:
        path: bash
        args:
        - -c
        - |2-

          set -eux
          echo "$HELLO_STR"
      params:
        HELLO_STR: ""
    params:
      HELLO_STR: hello-world
    attempts: 2
groups:
- name: a-group
  jobs:
  - regulator
```

## Random
* Why golang?  
Why not? Regardless the things I don't like about golang, I like it in general. Maybe Ruby is a better fit,
it does not provide strong static lint, which is one thing I want.

<!--* A language for Concourse only (like HCL)?-->

<!--Too much work.-->


## Note

Feedback / issue / pr welcome! I would also want to learn what is a good way to manage the pipelines.
