// Copyright (C) 2019 rameshvk. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"

	"github.com/google/go-cmp/cmp"
)

// Errorf is the type of the function used for reporting errors.
//
// This is typically testing.T.Error or testing.T.Fatal
type Errorf func(args ...interface{})

// Artifact implements comparing an artifact against a golden
// file.
//
// The goldenFile is expected to be relative to the testdata/ folder
// of the caller of this API.  The storage format is JSON for
// readability of the output.
//
// If the tests are run with -golden flag, the output is not compared
// but instead the output files are generated.
//
// Example Usage:
//
//    test.File(t.Fatal, "golden_xyz.json", func() interface{} {
//       ... do some tests and return any serializable type...
//    })
//
func Artifact(errorf Errorf, outputFile string, value interface{}) {
	var actual, expected interface{}

	pc := []uintptr{0}
	runtime.Callers(2, pc)
	f, _ := runtime.CallersFrames(pc).Next()

	outputFile = filepath.Join(filepath.Dir(f.File), "testdata/"+outputFile)

	bytes, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		errorf("Could not marshal value", err)
		return
	}

	if *goldenFlag {
		if err := ioutil.WriteFile(outputFile, bytes, 0644); err != nil {
			errorf("Could not save golden output", outputFile, err)
		}
		return
	}

	if err := json.Unmarshal(bytes, &actual); err != nil {
		errorf("Could not unmarshal value", err)
		return
	}

	if bytes, err = ioutil.ReadFile(outputFile); err != nil {
		errorf("error reading", outputFile, err)
		return
	}

	if err := json.Unmarshal(bytes, &expected); err != nil {
		errorf("could not unmarshal golden file", err)
		return
	}

	if diff := cmp.Diff(expected, actual); diff != "" {
		errorf("unexpected output", diff)
	}
}
