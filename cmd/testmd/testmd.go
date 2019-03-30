// Copyright (C) 2019 rameshvk. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

// Command testmd generates test files out of markdown snippets
//
//    $ go get github.com/tvastar/test/cmd/testmd
//
// This is useful when writing tutorials in markdown and making sure
// all the tutorials remain valid as the code underneath changes
//
// If no output file is specified (via -o), a temporary go file is
// generated and it is immediately run via `go test`.   Any
// additional arguments are passed through to go test.
//
// If a package name is provided and it does not include '_test', it
// will be assumed to be a runnable go program and "go run" will be
// used instead of "go test".
//
// The snippets of code must use the markdown code fence with either
// no info at all or an info string starting with "go"
//
// Additional fields on the info string can provide the name of the
// test:
// 
//     ```golang TestSomething
//     ....
//     ```
//
// If the name starts with anything but Test, an empty function with
// no args is produced which conveniently works for Examples
//
// Global variables can be introduced by using the special name of
// global:
//
//     ```golang global
//     var hidden := flag.Bool("hidden", false, "is it hidden?")
//     ```
//
// Snippets can be marked to be ignored:
//
//      ```golang skip
//      ...
//      ```
//
// A single markdown can be used to generate multiple package
// entries.  A snippet can be designated as being limited to a
// specific package by using the name of the package:
//
//      ```golang script_one.global
//      ....
//      ```
//
// The snippet above will only be included if the package name is
// script_one. This is useful when describing multiple main package
// options in one tutorial.
//
// Snippets can indicate import paths via a comment:
//
//       ```golang
//       // import fmt
//       ....
//       ```
//
// Usage:
//
//    $ testmd -pkg readme_test -o readme_test.go README.md
//    $ testmd README.md -v
//
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/tvastar/test"
)

//go:generate go run testmd.go -o testmd_test.go -pkg main_test ../../README.md

var output = flag.String("o", "", "output test file name")
var pkg = flag.String("pkg", "", "test package name")

func main() {
	flag.Parse()
	if flag.Arg(0) == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	pkgName := *pkg
	if pkgName == "" {
		pkgName = "test"
	}

	if *output != "" {
		fail(test.Markdown(flag.Arg(0), *output, pkgName))
		return
	}

	f, err := ioutil.TempFile("", "*_"+pkgName+".go")
	fail(err)
	name := f.Name()
	fail(f.Close())
	defer func() {
		fail(os.Remove(name))
	}()

	fail(test.Markdown(flag.Arg(0), name, pkgName))
	tool := "run"
	if strings.HasSuffix(pkgName, "test") {
		tool = "test"
	}
	args := append([]string{tool, name}, flag.Args()[1:]...)
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fail(cmd.Run())
}

func fail(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
