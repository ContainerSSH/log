package main

import (
	"fmt"
	"os"

	"github.com/containerssh/log"
)

func main() {
	source := "codes.go"
	destination := "CODES.md"

	if len(os.Args) > 1 {
		if len(os.Args) == 2 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
			println("Usage: containerssh-generate-codes [source.go DESTINATION.md]")
			os.Exit(0)
		} else if len(os.Args) != 3 {
			println("Usage: containerssh-generate-codes [source.go DESTINATION.md]")
			os.Exit(1)
		}
		source = os.Args[1]
		destination = os.Args[2]
	}

	log.MustWriteMessageCodesFile(source, destination)
	fmt.Printf("%s successfully written.\n", destination)
}
