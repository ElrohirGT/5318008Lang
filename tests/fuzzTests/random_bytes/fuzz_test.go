package fuzztests

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	lib "github.com/ElrohirGT/5318008Lang/applib"
)

const MIN_ITERS_PER_SEED = 100
const MAX_ITERS_PER_SEED = 1_000

const MIN_INPUT_SIZE = 50
const MAX_INPUT_SIZE = 1300

const MAX_BYTE = 256

const OUTPUT_SEPARATOR = "---"

func Fuzz_RandomInputStream(f *testing.F) {
	err := filepath.WalkDir("../../../tests/semantic_analysis/", func(path string, d fs.DirEntry, _ error) error {
		if d.IsDir() {
			return nil
		}
		contents, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		sContents := string(contents)
		parts := strings.Split(sContents, OUTPUT_SEPARATOR)
		f.Add(parts[0])
		return nil
	})
	if err != nil {
		panic(err)
	}

	f.Fuzz(func(t *testing.T, b string) {
		reader := bytes.NewBufferString(b)
		_ = lib.TestableMain(reader, lib.CompilerConfig{})
	})
}
