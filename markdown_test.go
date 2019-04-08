// Copyright (C) 2019 rameshvk. All rights reserved.
// Use of this source code is governed by a MIT-style license
// that can be found in the LICENSE file.

package test_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/tvastar/test"
)

func TestMarkdown(t *testing.T) {
	test.File(t.Fatal, "markdown.md", "markdown_test.go", func(src string) (string, error) {
		outf, err := ioutil.TempFile("", "*.md.test")
		if err != nil {
			return "", err
		}
		n := outf.Name()
		ignore(outf.Close())

		defer func() { ignore(os.Remove(n)) }()
		sources := []string{"testdata/markdown.md"}
		if err := test.Markdown(sources, n, "main"); err != nil {
			return "", err
		}

		data, err := ioutil.ReadFile(n)
		return string(data), err
	})
}

func TestMarkdownError(t *testing.T) {
	src := []string{"non_existent.md"}
	if err := test.Markdown(src, "boo", "main"); err == nil {
		t.Fatal("Did not fail on non-existent file")
	}
}

func ignore(err error) {}
