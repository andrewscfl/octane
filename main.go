package main

import (
	"flag"
	"fmt"
	definitions "octane/definitions"
	"os"
)

func main() {
	isRecursive := flag.Bool("r", false, "Recursive")
	target := os.Args[len(os.Args)-1]
	flag.Parse()
	fmt.Println("Recursive:", *isRecursive)
	fmt.Println("Target:", target)
	err := definitions.HandleDefinitionsLoading()
	if err != nil {
		panic(err)
	}

}
