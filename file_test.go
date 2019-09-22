// Copyright (C) 2019 rameshvk. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package test_test

import (
	"errors"
	"flag"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"

	"github.com/tvastar/test"
)

func identity(input string) (string, error) {
	return input, nil
}

var failError = errors.New("failure")

func fail(input string) (string, error) {
	return "", failError
}

func TestNoInput(t *testing.T) {
	failed := false
	errorf := func(args ...interface{}) {
		failed = true
	}

	test.File(errorf, "non-existent", "output.txt", identity)
	if !failed {
		t.Error("Failed to fail")
	}
}

func TestNoOutput(t *testing.T) {
	failed := false
	errorf := func(args ...interface{}) {
		failed = true
	}

	test.File(errorf, "input.txt", "non-existent", identity)
	if !failed {
		t.Error("Failed to fail")
	}
}

func TestSuccess(t *testing.T) {
	id1 := func(input string) string {
		return input
	}
	id2 := func(input []byte) []byte {
		return input
	}
	id3 := func(input []rune) []rune {
		return input
	}
	id4 := identity

	id5 := func(input struct{ OK int }) struct{ OK int } {
		return input
	}
	id6 := func(input *struct{ OK int }) struct{ OK int } {
		return *input
	}
	test.File(t.Fatal, "input.txt", "output.txt", id1)
	test.File(t.Fatal, "input.txt", "output.txt", id2)
	test.File(t.Fatal, "input.txt", "output.txt", id3)
	test.File(t.Fatal, "input.txt", "output.txt", id4)
	test.File(t.Fatal, "input.json", "output.json", id5)
	test.File(t.Fatal, "input.json", "output.json", id6)
}

func TestInvalidJSON(t *testing.T) {
	id := func(fn func()) string {
		return "boo"
	}
	var failed bool
	errorf := func(args ...interface{}) {
		failed = true
	}

	test.File(errorf, "input.json", "output.json", id)
	if !failed {
		t.Error("Unexpected success")
	}

	failed = false
	id2 := func(string) func() {
		return nil
	}
	test.File(errorf, "input.json", "output.json", id2)
	if !failed {
		t.Error("Unexpected success")
	}
}

func TestFunctionFail(t *testing.T) {
	var failure error
	errorf := func(args ...interface{}) {
		failure = args[0].(error)
	}

	test.File(errorf, "input.txt", "output.txt", fail)
	if failure != failError {
		t.Error("Unexpected failure", failure)
	}
}

func TestFunctionDiff(t *testing.T) {
	var diff string
	errorf := func(args ...interface{}) {
		diff = args[1].(string)
	}

	modify := func(input string) (string, error) {
		lines := strings.Split(input, "\n")
		return "fist of fury\n" + strings.Join(lines[1:], "\n"), nil
	}

	expected := `
- 	"first line",
+ 	"fist of fury",
  	"second line",`

	test.File(errorf, "input.txt", "output.txt", modify)
	if !strings.Contains(diff, expected) {
		t.Error("Unexpected diff", diff, "expected:", expected)
	}
}

func TestFileGolden(t *testing.T) {
	_, fname, _, _ := runtime.Caller(0)

	defer setGoldenFlag(false)
	setGoldenFlag(true)
	defer os.Remove(path.Join(path.Dir(fname), "testdata/golden.txt"))

	os.Remove(path.Join(path.Dir(fname), "testdata/golden.txt"))
	test.File(t.Error, "input.txt", "golden.txt", identity)
}

func TestFileGoldenWriteFail(t *testing.T) {
	defer setGoldenFlag(false)
	setGoldenFlag(true)

	var failed []interface{}
	errorf := func(args ...interface{}) {
		failed = args
	}
	test.File(errorf, "input.txt", "b/b", identity)
	if failed == nil {
		t.Error("Unexpected success writing to b/b")
	}
}

func setGoldenFlag(enable bool) {
	v := "false"
	if enable {
		v = "true"
	}
	if err := flag.Set("golden", v); err != nil {
		panic(err)
	}
}
