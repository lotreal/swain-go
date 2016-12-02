package main

import (
	"log"
	"os"
)

func main() {
	log.Printf("[INFO] %s \n", humanVersion)

	cli := NewCLI(os.Stdout, os.Stderr)
	os.Exit(cli.Run(os.Args))
}
