package main

import (
	"bytes"
	"fmt"
	"os"

	innerLib "github.com/ElrohirGT/5318008Lang/applib"
	lib "github.com/ElrohirGT/5318008Lang/lib"
)

func main() {
	filePath := os.Args[1]
	reader, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	err = innerLib.TestableMain(reader, innerLib.CompilerConfig{
		TACBuffer: lib.NewOpValue(bytes.Buffer{}),
		ASMBuffer: lib.NewOpValue(bytes.Buffer{}),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("No type errors found!")
}
