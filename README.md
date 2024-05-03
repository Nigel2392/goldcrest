Goldcrest- hooks
================

A simple hooks library for your Go projects.
Highly inspired from [Wagtail&#39;s](https://wagtail.org) hooks system.

## How to install

Easily install this with `go get`:

```go
go get github.com/Nigel2392/goldcrest@latest
```

## When to use

Hooks should be registered at initialization, this should probably not be done dynamically.

## How to use

We can register hooks by specyfing a name, order and functions for that hook.

It is recommended to also create a separate function type for your hook. This is not a requirement, but helps with documenting and overall readability.

Example:

```go
type MyIntFunc func([]int) []int

hooks.Register(
	"construct_int_list", -1,
	func(list []int) []int {
		return append(list, 1)
	},
)
```

To actually retrieve the results / execute the hooks we need to retrieve them from the registry.

```go
var myIntList = make([]int, 0)
myIntList = append(myIntList, 1, 2, 3)

var hookList = hooks.Get[MyIntFunc]("construct_int_list")
for _, hook := range hookList {
	myIntList = hook(myIntList)
}

```

We can also unregister the hook. 

**Be mindful!** It will delete all functions registered to it too.
