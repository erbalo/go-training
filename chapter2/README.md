# Chapter 2 - The looks

Chapter 2 - The looks

## Visibility

Go has a simple model for access control on function, methods, fields, etc.

If the name start with a capital letter then it is accessible from the outside (`exported`) else it is not (`unexported`). There is no other access level modifier available e.g. `internal`, `protected`.

Proper design is essential in order to work with this simple model.

## Documentation

Documentation is 1st class citizen in go. The default installation up to go version 1.12 contains a tool that shows the documentation of your current project via CLI or Web.

```bash
godoc -http=:6060
```

After the upgrade to go 1.13 you have to manually download the utility via:

```bash
go get golang.org/x/tools/cmd/godoc
godoc
```

The primary convention here is to comment anything that is exported because this is the public api of the package. You can see some details following the [link](https://golang.org/doc/effective_go.html#commentary).

The format of this comment should be as follows:

```go
// Sum calculates the summary of two integers.
// This is a second line comment.
func Sum(a, b int) int {
    return a + b
}
```

The comment always starts with a space and the name of the function and continues like reading a sentence. It always ends with a punctuation.

You can see more details following the [link](https://golang.org/doc/effective_go.html#commentary).

## Naming conventions

### Packages

- short, prefer transport over transportmechanism
- clear, like logging, postgres
- singular, e.g. `user` and not the plural `users`
- avoid catchall packages like utils, helpers, models
- since the package name is part of the declaration of a type we should use it e.g. a package named `cache` with a constructor `NewCache()` can be rename to just `New()` since the usage will always contain the package name `cache.New()`

### Variables

- Use `MixedCaps` or `mixedCaps` rather than underscores to write multi-word names.
- Abbreviations should always be capitalized e.g. `ServerHTTP`.
- single letter for indices (i, j, k in for loops)
- short names like "cust" for Customer or even "c" are perfectly valid as long as the declaration and its usages is very close. `The greater the distance between a name's declaration and its uses, the longer the name should be.`
- use repeated letters to represent a slice/array and use single letter in loops

```go
var tt []Thing

for i, t := range tt {
    ...
}
```

### Functions and methods

- Avoid repeating the package name in name of function and methods e.g. prefer `log.Info()` than `log.LogInfo()`.
- Go does not have getters and setters so the convention given a unexported field named `age` is `Age()` for the getter and `SetAge(age int)` for the setter.

### Interfaces

The name of the interface should be whatever the function is plus a "er" at the end

```go
type Reader interface {
    Read() ([]byte,error)
}
```

The above does not always make sense so try, but do not force it

```go
type Repository interface {
    Customer() (*Customer, error)
    Save(c *Customer) error
    Delete(id int) error
}
```

When embedding interfaces we should concatenate the name e.g. given a `Reader` and a `Closer` interface, the composition would be `ReadCloser`.

## Package structure

There is unfortunately no standard about structuring the packages. There are some guidelines like:

- [Standard Package Layout by Ben Johnson](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1), example [hasura](https://github.com/taxibeat/hasura)
- [Code like the Go team](https://www.youtube.com/watch?v=MzTcsI6tn-0)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout), example [REMS](https://github.com/taxibeat/rems), following the Ports and Adapters Architecture

Some of the above guidelines might be outdated or are not endorsed anymore.

## Project Structure

All project contain a lot of files for different purposes:

- Go files for our code
- Dockerfile
- docker-compose file
- CI files (Travis/CircleCI/Jenkins)
- Deployment files (Jenkins)
- Makefile
- Readme.md (root directory)
- .gitignore file
- go.mod, go.sum
- various scripts
- vendor
- Helm charts
- Infrastructure (terraform)
- Monitoring files (Grafana, Prometheus alerts)
- etc.

As you might see the list is big and we need to organize them a lot better in order to make it more maintainable.

We might organize our projects following the below structure (project named `myservice`):

- [myservice](#root-myservice)
  - [build [REQ]](#build)
    - [Dockerfile [REQ]](#dockerfile)
    - [Jenkinsfile.ci [OPT]](#jenkinsfile-ci)
  - [cmd [REQ]](#cmd)
  - [config [OPT]](#config)
  - [deployment [REQ]](#deployment)
    - [local [OPT]](#local)
    - [helm [OPT]](#helm)
    - [infra [OPT]](#infra)
    - [Jenkinsfile.cd [REQ]](#Jenkinsfile-cd)
  - [doc [REQ]](#doc)
  - [example [OPT]](#example)
  - [internal [REQ]](#internal)
  - [observability [REQ]](#observability)
    - [dashboard [REQ]](#dashboard)
    - [alerting [REQ]](#alerting)
  - [pkg [OPT]](#pkg)
  - [script [OPT]](#script)
  - [test [OPT]](#test)
  - [tool [OPT]](#tool)
  - [vendor [OPT]](#vendor)
  - [.travis.yml [OPT]](#travis)
  - [go.mod and go.sum [OPT]](#go-modules)
  - [Makefile [OPT]](#makefile)
  - [README.md [REQ]](#readme-file)
  - [.gitignore [REQ]](#gitignore)

Some of the items are optional [OPT] or required [REQ]. The above layout is heavily influenced by [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

Conventions here are the following:

- folder/package names are in singular format
- folder/package names should be short

### root (myservice)

The following files are contained, along with the other folders:

#### travis

Unfortunately, it has to be in the root.

#### go modules

These files are generated by the go modules tool.

#### Makefile

Makefile is explained in [chapter 5](../chapter5/README.md#makefile).

#### README file

Readme contains entry-point of the service documentation, which will be displayed in Github.

#### gitignore

The standard .gitignore file.

### build

Packaging and Continuous Integration. This folder should contains

#### Dockerfile

The standard dockerfile to create the deployment artifact of the service.

#### Jenkinsfile CI

If the project uses Jenkins as a CI server the ci file should be in here.

### cmd

Main applications for this project.  
The directory name for each application should match the name of the executable you want to have e.g. `/cmd/myservice/main.go`.

### config

Configuration file templates or default configs e.g. `env` files.

### deployment

Group deployment related files and folders.  
The root folder should contain:

- `Jenkinsfile.cd`

#### local

Local deployment setup e.g. `docker-compose`. Local Kubernetes will deprecate this.

#### helm

Helm packages for the deployment. Each deployable unit should have a sub-folder.

#### infra

Infrastructure as code e.g. `terraform` files.

#### Jenkinsfile CD

The Jenkinsfile responsible for deployment of the service.

### doc

Design and user documents (in addition to your godoc generated documentation).  
Things that could be in there are:

- OpenAPI documentation used by [Hypatia](https://github.com/taxibeat/hypatia)
- Architecture diagrams
- Run-books
- etc.

### example

Examples for your applications and/or public libraries.

### internal

Private application and library code that should not be imported by other projects.  

As an example of a project that follows [DDD](https://en.wikipedia.org/wiki/Domain-driven_design) and the [Ports and Adapters](https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)) Architecture approach the layout could contain the following sub-folders:

- app/domain, where all the domain code resides along with the interfaces [Ports] that communicate with infrastructure (DB, HTTP, e.g.)
- infra, where all concrete implementations [Adapters] of the above interfaces [Ports] reside e.g. [REMS](https://github.com/taxibeat/rems)

### observability

Artifacts needed to observe our service.

#### dashboard

Dashboards for Grafana.

#### alerting

Alerts for Prometheus AlertManager.

### pkg

Library code that's ok to use by external applications.

### script

Contains all scripts of the project.

### test

Additional external test apps and test data. Functional tests should live here.

### tool

Supporting tools for this project.

### vendor

Application dependencies.

## Formatting

Code formatting in go is actually a pretty simple thing. There is a tool that does almost all the formatting.
No need for policies, no need for holy wars.

One of the Go proverbs is:

**Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.**

The only thing you have to do is to call:

```bash
go fmt ./...
```

and every go file in every package will be formated with the only style that actually matters.
Visual Studio Code will setup this formatting automatically on each save and Goland has to be setup with file watchers.

The only thing that *gofmt* does not handle is line wrapping. Go has no line length limit. If a line feels too long, wrap it and indent with an extra tab.

[-> Next&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;: **Chapter 3**](../chapter3/README.md)  
[<- Previous&nbsp;: **Chapter 1**](../chapter1/README.md)
