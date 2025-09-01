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
	// "class_decl",
}

var IGNORE_SPECIFIC = []string{
	// "tests/semantic_analysis/typechecking/method_calling.cps_test",
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
		if err != nil && strings.TrimSpace(programOutput) != errMsg {

			// pos, ra, rb, _ := diffChar(programOutput, errMsg)
			// fmt.Printf("CHAR CIFF AT %d, %q vs %q \n%s\n%s\n", pos, ra, rb)

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

// diffChar returns the first differing rune and its position.
// If the strings are identical, ok == false.
func diffChar(a, b string) (pos int, ra, rb rune, ok bool) {
	ar, br := []rune(a), []rune(b)
	n := len(ar)
	if len(br) < n {
		n = len(br)
	}
	for i := 0; i < n; i++ {
		if ar[i] != br[i] {
			return i, ar[i], br[i], true
		}
	}
	// If all runes matched but lengths differ,
	// report the first extra rune.
	if len(ar) != len(br) {
		if len(ar) > len(br) {
			return n, ar[n], -1, true // -1 means "missing"
		}
		return n, -1, br[n], true
	}
	return 0, 0, 0, false
}
