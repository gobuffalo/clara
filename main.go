package main

import (
	"context"
	"log"
	"time"

	"github.com/gobuffalo/clara/v2/genny/rx"
	"github.com/gobuffalo/genny/v2"
	"github.com/gobuffalo/logger"
	"github.com/gobuffalo/meta"
	"github.com/spf13/pflag"
)

var options = struct {
	*rx.Options
	dryRun  bool
	verbose bool
}{
	Options: &rx.Options{},
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	flags := pflag.NewFlagSet("clara", pflag.ExitOnError)

	flags.BoolVarP(&options.dryRun, "dry-run", "d", false, "dry run")
	flags.BoolVarP(&options.verbose, "verbose", "v", false, "verbose output")
	flags.BoolVar(&options.SkipBuffalo, "skip-buffalo", false, "skip buffalo related checks")
	flags.BoolVar(&options.SkipNode, "skip-node", false, "skip node related checks")
	flags.BoolVar(&options.SkipDB, "skip-db", false, "skip DB related checks")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	run := genny.WetRunner(ctx)
	if options.dryRun {
		run = genny.DryRunner(ctx)
	}

	if options.verbose {
		run.Logger = logger.New(logger.DebugLevel)
	}

	opts := options.Options
	opts.App = meta.New(".")
	if err := run.WithNew(rx.New(opts)); err != nil {
		return err
	}
	return run.Run()
}
