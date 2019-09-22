# Test

[![Status](https://travis-ci.com/tvastar/test.svg?branch=master)](https://travis-ci.com/tvastar/test?branch=master)
[![GoDoc](https://godoc.org/github.com/tvastar/test?status.svg)](https://godoc.org/github.com/tvastar/test)
[![codecov](https://codecov.io/gh/tvastar/test/branch/master/graph/badge.svg)](https://codecov.io/gh/tvastar/test)
[![Go Report Card](https://goreportcard.com/badge/github.com/tvastar/test)](https://goreportcard.com/report/github.com/tvastar/test)

Test is a simple golang test utils package.

## test.Artifact

test.Artifact allows a simple way to test artifacts against recorded
golden files.

The golden files are stored in `testdata/` folder of the caller of the
API and are expected to contain the JSON version of the value being
tested.  The [go-cmp](https://github.com/google/go-cmp) package is
used for better quality diffs.

## test.File

This is deprecated in favor of test.Artifact.

## test.Markdown

Markdown() converts a set of code snippets in markdown into test cases. For example:

````
```go
if len("hello") != 5 {
	fmt.Println("This should not happen")
}

// Output:
```
````

The above snippet is converted to an `ExampleXXX()` which runs the code.

Snippets can also be named, with ability to import code from elsewhere:

````
```go  TestBoo
// import runtime
_, fname, _, _ := runtime.Caller(1)
if fname == "" {
	t.Error("Could not find caller info")
}
```
````

The name of the snippet must be TestXYZ or ExampleXYZ or just a plain
function (presumably for use in other snippets).  Two sepcial cases
are "skip" to indicate the fence block should not be considered
testable and "global" to indicate the fence block is a global
variable:

````
```go skip
   // this has sample code that is not expected to compile
   booya
```
````
````
```go global
func stranger() string {
	return "stranger"
}
```
````
````
```go Example_UsingSomethingDefinedInGlobal
fmt.Println("ok", stranger())

// Output: ok stranger
```
````
If the info name has a dot in it, the word before the dot is compared
to the package name. If they don't match the fence block is
skipped. This allows a single markdown to have multiple fence blocks
destined for different packages:

````
```go firstPackage.DoSomething
...
```

```go secondPackage.
...
```
````
