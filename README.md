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
