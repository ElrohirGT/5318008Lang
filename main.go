package main

import (
	"fmt"
	"os"

	lib "github.com/ElrohirGT/5318008Lang/applib"
)

func main() {
	filePath := os.Args[1]
	reader, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	err = lib.TestableMain(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println("No type errors found!")
}
