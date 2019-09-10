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
