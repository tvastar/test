# Test

[![Status](https://travis-ci.com/tvastar/test.svg?branch=master)](https://travis-ci.com/tvastar/test?branch=master)
[![GoDoc](https://godoc.org/github.com/tvastar/test?status.svg)](https://godoc.org/github.com/tvastar/test)
[![codecov](https://codecov.io/gh/tvastar/test/branch/master/graph/badge.svg)](https://codecov.io/gh/tvastar/test)
[![Go Report Card](https://goreportcard.com/badge/github.com/tvastar/test)](https://goreportcard.com/report/github.com/tvastar/test)

Test is a simple golang test utils package.

## test.File

test.File implements a simple file testing utility that:

1. reads input file in the testdata/ directory into a string
2. runs this against the provided test function
3. compares against the contents of the output file (also in the testdata/ directory) 
4. reports discrepancies in the diff format (such as the one used by git)

It also accepts a -golden flag which when specified skips the
comparison and instead saves the output provided by the function into
the output file.  This makes creation of initial output files and
updates relatively easy

## test.Markdown

Markdown() converts a set of code snippets in markdown into test cases. For example:

```go

if len("hello") != 5 {
	fmt.Println("This should not happen")
}

// Output:
```

The above snippet is converted to an `ExampleXXX()` which runs the code.

Snippets can also be named, with ability to import code from elsewhere:

```go  TestBoo

// import runtime
_, fname, _, _ := runtime.Caller(1)
if fname == "" {
   t.Error("Could not find caller info")
}
```

The name of the snippet must be TestXYZ or ExampleXYZ or just a plain
function (presumably for use in other snippets).  Two sepcial cases
are "skip" to indicate the fence block should not be considered
testable and "global" to indicate the fence block is a global
variable:

```go skip
   // this has sample code that is not expected to compile
   booya
```

```go global
func stranger() string {
	return "stranger"
}
```

```go Example_UsingSomethingDefinedInGlobal

fmt.Println("ok", stranger())

// Output: ok stranger
```





