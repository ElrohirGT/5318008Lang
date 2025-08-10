package main

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const OUTPUT_SEPARATOR = "---"

var RUN_ONLY_THAT_MATCH = []string{
	"basic_expre",
}

func Test_SnapshotTesting(t *testing.T) {
	filePaths := []string{}
	filepath.WalkDir("./tests/", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		for _, str := range RUN_ONLY_THAT_MATCH {
			if !strings.Contains(path, str) {
				return nil
			}
		}

		filePaths = append(filePaths, path)
		return nil
	})

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
		if err != nil && programOutput != strings.TrimSpace(err.Error()) {
			t.Errorf(
				"\nProgram %s failed with:\n%s\nBut expected:\n%s",
				path,
				err.Error(),
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
