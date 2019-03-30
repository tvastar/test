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

```go main.global
func stranger() string {
	return "stranger"
}
```

```go Example_UsingSomethingDefinedInGlobal

fmt.Println("ok", stranger())

// Output: ok stranger
```

```go unused.Garbage
This will be skipped because of package name
```

```python
python scripts should get dropped
```
