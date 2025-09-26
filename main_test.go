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

	lib "github.com/ElrohirGT/5318008Lang/applib"
)

const OUTPUT_SEPARATOR = "---"

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}

var RUN_ONLY_THAT_MATCH = []string{
	// "basic_expre",
	// "class",
	// "tests/semantic_analysis/controlflow/",
	// "tests/semantic_analysis/scopes/break_outside_loop",
	// "typechecking",
	// "class_constructor",
	// "simplified.cps_test",
}

var IGNORE_SPECIFIC = []string{
	// "tests/semantic_analysis/typechecking/method_calling.cps_test",
	// "tests/semantic_analysis/typechecking/class_chaining.cps_test",
}

func Test_SemanticAnalysis(t *testing.T) {
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
		expectedOutput := strings.TrimSpace(parts[1])

		reader := bytes.NewReader([]byte(cpsContents))
		err = lib.TestableMain(reader)
		errMsg := ""
		if err != nil {
			errMsg = strings.TrimSpace(stripANSI(err.Error()))
		}

		if expectedOutput != errMsg {
			b := strings.Builder{}
			b.WriteString("\nProgram ")
			b.WriteString(path)
			b.WriteString(" failed with:\n")

			lastI := 0
			for i, expectedByte := range []byte(expectedOutput) {
				if i < len(errMsg) {
					actualByte := errMsg[i]
					if actualByte != expectedByte {
						b.WriteString(lib.Red)
					}
					b.WriteByte(actualByte)
					b.WriteString(lib.Reset)
				} else {
					b.WriteString(lib.Grey)
					b.WriteByte(expectedByte)
					b.WriteString(lib.Reset)
				}
				lastI = i
			}

			if lastI+1 < len(errMsg)-1 {
				b.WriteString(lib.Red)
				b.WriteString(errMsg[lastI+1:])
				b.WriteString(lib.Reset)
			}

			b.WriteString("\nBut expected:\n")
			b.WriteString(expectedOutput)

			t.Error(b.String())
			continue
		}
	}
}
