package main

import (
	"flag"
	"fmt"
	"os"
)

var nameFlagVar string

func main() {
	flag.StringVar(&nameFlagVar, "n", "", "set your name (shorthand)")
	flag.Parse()
	flag.Usage = func() {
		fmt.Println("Usage: greet -n \"John\"")
	}

	if len(os.Args) <= 1 {
		flag.Usage()
		return
	}

	switch os.Args[1] {
	case "-n":
		fmt.Printf("Hello, %s\n", nameFlagVar)
	default:
		flag.Usage()
		os.Exit(1)
	}

}
