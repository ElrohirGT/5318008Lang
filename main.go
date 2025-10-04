package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	innerLib "github.com/ElrohirGT/5318008Lang/applib"
	lib "github.com/ElrohirGT/5318008Lang/lib"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "filePath", os.Args[1], "File path of the source code")

	var outPath string
	flag.StringVar(&outPath, "o", "out.asm", "File path to the final executable")

	var tacPath string
	flag.StringVar(&tacPath, "tac", "out.ir", "File path to the intermediate representation")
	flag.Parse()

	reader, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	tacBuffer := bytes.Buffer{}
	asmBuffer := bytes.Buffer{}

	err = innerLib.TestableMain(reader, innerLib.CompilerConfig{
		TACBuffer: lib.NewOpValue(&tacBuffer),
		ASMBuffer: lib.NewOpValue(&asmBuffer),
	})
	if err != nil {
		panic("\n" + err.Error())
	}

	err = os.WriteFile(tacPath, tacBuffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(outPath, asmBuffer.Bytes(), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("No type errors found!")
}
