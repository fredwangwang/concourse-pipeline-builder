# Concourse Pipeline Builder

This is a kinda POC thing to make the pipelines more structural.

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
- [ ] Make it a cli?
- [ ] todo?

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
- name: pcf-pipelines
  type: git
  source:
    branch: master
    uri: git@github.com:pivotal-cf/pcf-pipelines.git
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
```

## Random
* Why golang?

Why not? Regardless the things I don't like about golang, I like it in general. Maybe Ruby is a better fit,
it does not provide strong static lint, which is one thing I want.

<!--* A language for Concourse only (like HCL)?-->

<!--Too much work.-->


## Note

Feedbacks welcome! I would also want to learn what is a good way to manage the pipelines.
