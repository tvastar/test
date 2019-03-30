// Copyright (C) 2019 rameshvk. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

// Command testmd generates test files out of markdown snippets
//
//    $ go get github.com/tvastar/test/cmd/testmd
//
// If no output file is specified (via -o), a temporary go file is
// generated and it is immediately tested via `go test`.   Any
// additional arguments are passed through to go test.
//
// Usage:
//
//    $ testmd README.md -pkg readme_test -o readme_test.go
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
