// Copyright (C) 2019 rameshvk. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

// Package test implements simple test utilities.
package test

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"runtime"
	"time"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// Errorf is the type of the function used for reporting errors.
//
// This is typically testing.T.Error or testing.T.Fatal
type Errorf func(args ...interface{})

// File implements testing against input/output files
//
// The input and output file names are relative to the testdata/
// folder of the caller.
//
// The provided function is one of these two forms:
//
//    func (input someType) (output someOtherType)
//    func (input someType) (output someOtherType,  err error)
//
// The input file is read and the contents passed through this
// function. For input arguments of type string, []byte or []rune the
// contents of the files are passed as is. For other types,  the
// contents are assumed to be JSON encoded.  The output is similarly
// JSON encoded for such types.
//
// The discrepancies are reported using regular diff format via the
// error function (which sports the same signature as testing.T.Error
// or testing.T.Fatal)
//
// If the tests are run with -golden flag, the output is not compared
// but instead the output files are created to match the output
// provided by the test.
//
// Example Usage:
//
//    test.File(t.Fatal, "input.txt", "output.txt",
//       func(input string) string { .... },
//    )
//
func File(errorf Errorf, inputFile string, outputFile string, fn interface{}) {
	pc := []uintptr{0}
	runtime.Callers(2, pc)
	f, _ := runtime.CallersFrames(pc).Next()

	inputFile = filepath.Join(filepath.Dir(f.File), "testdata/"+inputFile)
	outputFile = filepath.Join(filepath.Dir(f.File), "testdata/"+outputFile)

	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		errorf("error reading", inputFile, err)
		return
	}

	output, err := invoke(fn, string(bytes))
	if err != nil {
		errorf(err)
		return
	}

	if *goldenFlag {
		if err := ioutil.WriteFile(outputFile, []byte(output), 0644); err != nil {
			errorf("Could not save golden output", outputFile, err)
		}
		return
	}

	bytes, err = ioutil.ReadFile(outputFile)
	if err != nil {
		errorf("error reading", outputFile, err)
		return
	}

	if output != string(bytes) {
		errorf("unexpected output", diff(string(bytes), output))
	}
}

func diff(expected, got string) string {
	dmp := diffmatchpatch.New()
	dmp.DiffTimeout = time.Minute * 2
	wSrc, wDst, warray := dmp.DiffLinesToRunes(expected, got)
	diffs := dmp.DiffMainRunes(wSrc, wDst, false)
	diffs = dmp.DiffCharsToLines(diffs, warray)

	result := ""
	for _, line := range diffs {
		text := line.Text
		if text != "" && text[len(text)-1] == '\n' {
			text = text[:len(text)-1]
		}

		if line.Type == diffmatchpatch.DiffInsert {
			result += "\n+ " + text
		} else if line.Type == diffmatchpatch.DiffDelete {
			result += "\n- " + text
		}
	}
	return result
}

func invoke(fn interface{}, input string) (string, error) {
	v := reflect.ValueOf(fn)
	argType := v.Type().In(0)
	var results []reflect.Value

	switch reflect.Zero(argType).Interface().(type) {
	case string:
		results = v.Call([]reflect.Value{reflect.ValueOf(input)})
	case []byte:
		bytes := []byte(input)
		results = v.Call([]reflect.Value{reflect.ValueOf(bytes)})
	case []rune:
		runes := []rune(input)
		results = v.Call([]reflect.Value{reflect.ValueOf(runes)})
	default:
		ptr := reflect.New(argType)
		if err := json.Unmarshal([]byte(input), ptr.Interface()); err != nil {
			return "", err
		}
		results = v.Call([]reflect.Value{ptr.Elem()})
	}

	if len(results) > 1 {
		if err, _ := results[1].Interface().(error); err != nil {
			return "", err
		}
	}

	switch r := results[0].Interface().(type) {
	case string:
		return r, nil
	case []byte:
		return string(r), nil
	case []rune:
		return string(r), nil
	}

	bytes, err := json.MarshalIndent(results[0].Interface(), "", "\t")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

var goldenFlag = flag.Bool("golden", false, "build golden testdata files instead of verifying")
