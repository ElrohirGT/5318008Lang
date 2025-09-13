package fuzztests

import (
	"bytes"
	"io"
	"math/rand"
	"testing"

	lib "github.com/ElrohirGT/5318008Lang/applib"
)

const MIN_ITERS_PER_SEED = 1_000
const MAX_ITERS_PER_SEED = 5_000

const MIN_INPUT_SIZE = 2500
const MAX_INPUT_SIZE = 7000

const MAX_BYTE = 256

func generateRandomBytesReader(r *rand.Rand) io.Reader {
	size := r.Intn(MAX_ITERS_PER_SEED-MIN_ITERS_PER_SEED) + MIN_ITERS_PER_SEED
	b := make([]byte, 0, size)
	for range size {
		b = append(b, byte(r.Int()%MAX_BYTE))
	}
	return bytes.NewReader(b)
}

func Fuzz_RandomInputStream(f *testing.F) {
	f.Add(int64(4206969))
	f.Fuzz(func(t *testing.T, a int64) {
		randSource := rand.NewSource(a)
		r := rand.New(randSource)

		for range r.Intn(MAX_ITERS_PER_SEED-MIN_ITERS_PER_SEED) + MIN_ITERS_PER_SEED {
			reader := generateRandomBytesReader(r)
			lib.TestableMain(reader)
		}
	})
}
