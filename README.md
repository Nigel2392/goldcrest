Goldcrest- hooks
================

A simple hooks library for your Go projects.
Highly inspired from [Wagtail&#39;s](https://wagtail.org) hooks system.

## How to install

Easily install this with `go get`:

```go
go get github.com/Nigel2392/goldcrest@latest
```

## How to use

We can register hooks by specyfing a name, order and functions for that hook.

It is recommended to also create a separate function type for your hook. This is not a requirement, but helps with documenting and overall readability.

Example:

```go
type MyIntFunc func([]int) []int

goldcrest.Register(
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

var hookList = goldcrest.Get[MyIntFunc]("construct_int_list")
for _, hook := range hookList {
	myIntList = hook(myIntList)
}

```

We can also unregister the hook.

**Be mindful!** It will delete all functions registered to it too.

```go
goldcrest.Unregister("construct_int_list")
```

## Implementation Details

Hooks should be registered at initialization, this should probably not be done dynamically.

The registry is implemented as a map. *This means you will not be able to concurrently add hooks.*

We do however provide a way to create a separate registry. This will allow for more dynamic additions to the registry.

Example:

```go
// Create a custom hooks registry
var myHooksRegistry = make(goldcrest.HookRegistry)

// Register a hook to your custom registry
myHooksRegistry.register(
	// ... Same as before
)

// Get a hook from your custom registry.
// Note: we do not adress the custom registry directly.
var hookList = goldcrest.GetFrom[MyIntFunc](myHooksRegistry, "construct_int_list")
// ...

// Unregister a hook from your registry
myHooksRegistry.Unregister("construct_int_list")
```
