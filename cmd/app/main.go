package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/autobrr/crossplatform-playground/application"
)

var (
	version = "v0.0.0-dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

const usage = `
Usage: autobrr <action> [options]

Actions:
  version  Print version information
`

func init() {
	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), usage)
	}
}

func main() {
	flag.Parse()

	defer logPanic()

	switch cmd := flag.Arg(0); cmd {
	case "version":
		fmt.Printf("Version: %s\nCommit: %s\nDate: %s\nBuilt by: %s\n", version, commit, date, builtBy)

	default:
		app := application.NewApplication()
		ctx := context.Background()

		if err := app.Start(ctx); err != nil {
			os.Exit(1)
		}
	}
}

func logPanic() {
	if r := recover(); r != nil {
		fmt.Printf("Panic %v\n%s\n", r, string(debug.Stack()))
	}
}
