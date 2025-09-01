package main

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"
)

const OUTPUT_SEPARATOR = "---"

var RUN_ONLY_THAT_MATCH = []string{
	// "basic_expre",
	// "class",
	// "tests/semantic_analysis/controlflow/",
	// "tests/semantic_analysis/scopes/break_outside_loop",
	// "typechecking",
}

var IGNORE_SPECIFIC = []string{
	"tests/semantic_analysis/typechecking/method_calling.cps_test",
	"tests/semantic_analysis/typechecking/arrays.cps_test",
}

func Test_SnapshotTesting(t *testing.T) {
	filePaths := []string{}
	err := filepath.WalkDir("./tests/semantic_analysis", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		for _, str := range RUN_ONLY_THAT_MATCH {
			if !strings.Contains(path, str) {
				return nil
			}
		}

		if slices.Contains(IGNORE_SPECIFIC, path) {
			return nil
		}

		filePaths = append(filePaths, path)
		return nil
	})

	if err != nil {
		t.Error(err)
	}

	for _, path := range filePaths {
		fileBytes, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("\nFAILED %s:\n %s", path, err)
			continue
		}

		parts := strings.Split(string(fileBytes), OUTPUT_SEPARATOR)
		cpsContents := strings.TrimSpace(parts[0])
		programOutput := strings.TrimSpace(parts[1])

		reader := bytes.NewReader([]byte(cpsContents))
		err = testableMain(reader)
		errMsg := ""
		if err != nil {
			errMsg = strings.TrimSpace(stripANSI(err.Error()))
		}
		if err != nil && programOutput != errMsg {
			t.Errorf(
				"\nProgram %s failed with:\n%s\nBut expected:\n%s",
				path,
				errMsg,
				programOutput,
			)
			continue
		}
		if err == nil && programOutput != "" {
			t.Errorf(
				"\nProgram %s didn't fail!\nBut expected:\n%s",
				path,
				programOutput,
			)
			continue
		}
	}
}

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}
