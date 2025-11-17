package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"

	"github.com/ElrohirGT/5318008Lang/applib"
	lib "github.com/ElrohirGT/5318008Lang/lib"
)

const OUTPUT_SEPARATOR = "---"

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}

const ENABLE_PANIC_RECOV = false

// NOTE: Only change this flag when you're ABSOLUTELY CERTAIN all tests are correct
// but changing each test manually would take a bunch of time.
// MOST OF THE TIME THIS SHOULD BE FALSE!!!
const REGENERATE_TESTS = false

var RUN_ONLY_THAT_MATCH = []string{
	// "basic_expre",
	// "class",
	// "tests/semantic_analysis/controlflow/",
	// "tests/semantic_analysis/scopes/break_outside_loop",
	// "typechecking",
	// "class_constructor",
	// "TAC_generation/",
	// "code_generation/basic_control_flow",
}

var IGNORE_SPECIFIC = []string{
	// "tests/semantic_analysis/typechecking/method_calling.cps_test",
	// "tests/semantic_analysis/typechecking/class_chaining.cps_test",
	// "tests/code_generation/booleans.cps_test",
	"tests/code_generation/control_flow_with_functions.cps_test",
	"tests/code_generation/recursive.cps_test",
	"tests/code_generation/functions.cps_test",
	// "tests/code_generation/basic_control_flow.cps_test",
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
		t.Log("Testing:", path)
		err = applib.TestableMain(reader, applib.CompilerConfig{})
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
						b.WriteString(applib.Red)
					}
					b.WriteByte(actualByte)
					b.WriteString(applib.Reset)
				} else {
					b.WriteString(applib.Grey)
					b.WriteByte(expectedByte)
					b.WriteString(applib.Reset)
				}
				lastI = i
			}

			if lastI+1 < len(errMsg)-1 {
				b.WriteString(applib.Red)
				b.WriteString(errMsg[lastI+1:])
				b.WriteString(applib.Reset)
			}

			b.WriteString("\nBut expected:\n")
			b.WriteString(expectedOutput)

			t.Error(b.String())
			continue
		}
	}
}

func Test_TACGeneration(t *testing.T) {
	filePaths := []string{}
	err := filepath.WalkDir("./tests/TAC_generation/", func(path string, d fs.DirEntry, err error) error {
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
		t.Log("Reading file:", path)

		parts := strings.Split(string(fileBytes), OUTPUT_SEPARATOR)
		cpsContents := strings.TrimSpace(parts[0])
		expectedOutput := strings.TrimSpace(parts[1])

		outBuffer := bytes.Buffer{}
		reader := bytes.NewReader([]byte(cpsContents))

		if ENABLE_PANIC_RECOV {
			err = func() error {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("THE CODE CONTAINS A SKILL ISSUE!\nPanic: %s", r)
					}
				}()
				return applib.TestableMain(reader, applib.CompilerConfig{
					TACBuffer: lib.NewOpValue(&outBuffer),
				})
			}()
		} else {
			err = applib.TestableMain(reader, applib.CompilerConfig{
				TACBuffer: lib.NewOpValue(&outBuffer),
			})
		}

		actualOutput := strings.TrimSpace(outBuffer.String())
		if err != nil {
			actualOutput = strings.TrimSpace(stripANSI(err.Error()))
		}

		if REGENERATE_TESTS {
			// fileBytes, err := os.ReadFile(path)
			var b bytes.Buffer
			_, err := fmt.Fprintf(&b, "%s\n%s\n%s", cpsContents, OUTPUT_SEPARATOR, actualOutput)
			if err != nil {
				log.Panicf("Failed to generate file contents for: %s", path)
			}

			err = os.WriteFile(path, b.Bytes(), 0644)
			if err != nil {
				log.Panicf("Failed to regenerate test: %s", path)
			}
		}

		if expectedOutput != actualOutput {
			b := strings.Builder{}
			b.WriteString("\nProgram ")
			b.WriteString(path)
			b.WriteString(" failed with:\n")

			lastI := 0
			for i, expectedByte := range []byte(expectedOutput) {
				if i < len(actualOutput) {
					actualByte := actualOutput[i]
					if actualByte != expectedByte {
						b.WriteString(applib.Red)
					}
					b.WriteByte(actualByte)
					b.WriteString(applib.Reset)
				} else {
					b.WriteString(applib.Grey)
					b.WriteByte(expectedByte)
					b.WriteString(applib.Reset)
				}
				lastI = i
			}

			if lastI+1 < len(actualOutput)-1 {
				b.WriteString(applib.Red)
				b.WriteString(actualOutput[lastI+1:])
				b.WriteString(applib.Reset)
			}

			b.WriteString("\nBut expected:\n")
			b.WriteString(expectedOutput)

			t.Error(b.String())
			continue
		}
	}
}

func Test_ASMGeneration(t *testing.T) {
	filePaths := []string{}
	err := filepath.WalkDir("./tests/code_generation/", func(path string, d fs.DirEntry, err error) error {
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
		t.Log("Reading file:", path)

		parts := strings.Split(string(fileBytes), OUTPUT_SEPARATOR)
		cpsContents := strings.TrimSpace(parts[0])
		expectedOutput := strings.TrimSpace(parts[1])

		tacBuffer := bytes.Buffer{}
		asmBuffer := bytes.Buffer{}
		reader := bytes.NewReader([]byte(cpsContents))

		if ENABLE_PANIC_RECOV {
			err = func() error {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("THE CODE CONTAINS A SKILL ISSUE!\nPanic: %s", r)
					}
				}()
				return applib.TestableMain(reader, applib.CompilerConfig{
					TACBuffer: lib.NewOpValue(&tacBuffer),
					ASMBuffer: lib.NewOpValue(&asmBuffer),
				})
			}()
		} else {
			err = applib.TestableMain(reader, applib.CompilerConfig{
				TACBuffer: lib.NewOpValue(&tacBuffer),
				ASMBuffer: lib.NewOpValue(&asmBuffer),
			})
		}

		actualOutput := strings.TrimSpace(asmBuffer.String())
		if err != nil {
			actualOutput = strings.TrimSpace(stripANSI(err.Error()))
		}

		if REGENERATE_TESTS {
			// fileBytes, err := os.ReadFile(path)
			var b bytes.Buffer
			_, err := fmt.Fprintf(&b, "%s\n%s\n%s", cpsContents, OUTPUT_SEPARATOR, actualOutput)
			if err != nil {
				log.Panicf("Failed to generate file contents for: %s", path)
			}

			err = os.WriteFile(path, b.Bytes(), 0644)
			if err != nil {
				log.Panicf("Failed to regenerate test: %s", path)
			}
		}

		if expectedOutput != actualOutput {

			b := strings.Builder{}
			b.WriteString("\nProgram ")
			b.WriteString(path)
			b.WriteString(" failed with:\n")

			lastI := 0
			for i, expectedByte := range []byte(expectedOutput) {
				if i < len(actualOutput) {
					actualByte := actualOutput[i]
					if actualByte != expectedByte {
						b.WriteString(applib.Red)
					}
					b.WriteByte(actualByte)
					b.WriteString(applib.Reset)
				} else {
					b.WriteString(applib.Grey)
					b.WriteByte(expectedByte)
					b.WriteString(applib.Reset)
				}
				lastI = i
			}

			if lastI+1 < len(actualOutput)-1 {
				b.WriteString(applib.Red)
				b.WriteString(actualOutput[lastI+1:])
				b.WriteString(applib.Reset)
			}

			b.WriteString("\nBut expected:\n")
			b.WriteString(expectedOutput)

			b.WriteString("\nTAC:\n")
			b.WriteString(tacBuffer.String())

			t.Error(b.String())
			continue
		}
	}
}
