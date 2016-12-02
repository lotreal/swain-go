package main

import (
	// "errors"
	"flag"
	"io"
	"log"
	"os"
	"sync"
)

// CLI is the main entry point.
type CLI struct {
	sync.Mutex

	// outSteam and errStream are the standard out and standard error streams to
	// write messages from the CLI.
	outStream, errStream io.Writer

	// signalCh is the channel where the cli receives signals.
	signalCh chan os.Signal

	// stopCh is an internal channel used to trigger a shutdown of the CLI.
	stopCh  chan struct{}
	stopped bool
}

// NewCLI creates a new CLI object with the given stdout and stderr streams.
func NewCLI(out, err io.Writer) *CLI {
	return &CLI{
		outStream: out,
		errStream: err,
		signalCh:  make(chan os.Signal, 1),
		stopCh:    make(chan struct{}),
	}
}

// Run accepts a slice of arguments and returns an int representing the exit
// status from the command.
func (cli *CLI) Run(args []string) int {
	cli.Lock()
	defer cli.Unlock()

	cli.ParseFlags(args[1:])
	log.Printf("%#v\n", cli)
	log.Printf("%#v\n", args)
	return 0
}

func (cli *CLI) ParseFlags(args []string) {
	// Parse the flags and options
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	log.Printf("%#v\n", flags)
	log.Printf("%T\n", flags)
	flags.SetOutput(cli.errStream)

	log.Printf("%T\n", (funcVar)(func(s string) error {
		log.Printf("%#v\n", s)
		// return errors.New("error decoding config")
		return nil
	}))

	flags.Var((funcVar)(func(s string) error {
		log.Printf("%#v\n", s)
		// return errors.New("error decoding config")
		return nil
	}), "consul", "")

	flags.Parse(args)

	// Error if extra arguments are present
	args = flags.Args()
	if len(args) > 0 {
		log.Printf("cli: extra args: %q", args)
	}
}
