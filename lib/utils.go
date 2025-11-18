package lib

const MIPS32_WORD_BYTE_SIZE uint = 4

type Number interface {
	int | uint
}

func AlignSize[T Number](value T, wordSize T) T {
	var zero T

	if value%wordSize != zero {
		return value + (wordSize - (value % wordSize))
	}

	return value
}
