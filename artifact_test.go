// Copyright (C) 2019 rameshvk. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package test_test

import (
	"flag"
	"os"
	"path"
	"regexp"
	"runtime"
	"testing"

	"github.com/tvastar/test"
)

func TestArtifactNoOutput(t *testing.T) {
	failed := false
	errorf := func(args ...interface{}) {
		failed = true
	}

	test.Artifact(errorf, "non-existent", "boo")
	if !failed {
		t.Error("Failed to fail")
	}
}

func TestArtifactNotSerializable(t *testing.T) {
	failed := false
	errorf := func(args ...interface{}) {
		failed = true
	}

	test.Artifact(errorf, "non-existent", func() {})
	if !failed {
		t.Error("Failed to fail")
	}
}

func TestArtifactUnmarshalableGoldenFile(t *testing.T) {
	failed := false
	errorf := func(args ...interface{}) {
		failed = true
	}

	test.Artifact(errorf, "golden.txt", func() {})
	if !failed {
		t.Error("Failed to fail")
	}
}

func TestArtifactSuccess(t *testing.T) {
	defer restoreGoldenFlag()()

	_, fname, _, _ := runtime.Caller(0)
	defer os.Remove(path.Join(path.Dir(fname), "testdata/golden_artifact.json"))
	os.Remove(path.Join(path.Dir(fname), "testdata/golden_artifact.json"))

	value := map[string]string{"hello": "world", "boo": "hoo"}

	check(flag.Set("golden", "true"))
	test.Artifact(t.Error, "golden_artifact.json", value)
	check(flag.Set("golden", "false"))
	test.Artifact(t.Error, "golden_artifact.json", value)

	// also validate how it prints differences
	errorf := func(args ...interface{}) {
		expected := regexp.MustCompile(
			`-.*map\[string\]interface\{\}\{"boo": string\("hoo"\), "hello": string\("world"\)\}`,
		)

		if !expected.MatchString(args[1].(string)) {
			t.Error("Unexpected output", args[1])
		}
	}
	test.Artifact(errorf, "golden_artifact.json", nil)
}

func TestArtifactGoldenWriteFail(t *testing.T) {
	defer restoreGoldenFlag()()

	check(flag.Set("golden", "true"))

	var failed []interface{}
	errorf := func(args ...interface{}) {
		failed = args
	}
	test.Artifact(errorf, "b/b", "")
	if failed == nil {
		t.Error("Unexpected success writing to b/b")
	}
}

func restoreGoldenFlag() func() {
	before := flag.Lookup("golden").Value.String()
	return func() {
		check(flag.Set("golden", before))
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
