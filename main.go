package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Printf("start\n")
	fmt.Printf("[INFO] %s \n", humanVersion)

	cli := NewCLI(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args))
}
