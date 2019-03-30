package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gobuffalo/clara/genny/rx"
	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/logger"
	"github.com/gobuffalo/meta"
	"github.com/spf13/cobra"
)

var options = struct {
	*rx.Options
	dryRun  bool
	verbose bool
}{
	Options: &rx.Options{},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clara",
	Short: "A brief description of your application",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
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
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&options.dryRun, "dry-run", "d", false, "dry run")
	rootCmd.Flags().BoolVarP(&options.verbose, "verbose", "v", false, "verbose output")
	rootCmd.Flags().BoolVar(&options.SkipBuffalo, "skip-buffalo", false, "skip buffalo related checks")
	rootCmd.Flags().BoolVar(&options.SkipNode, "skip-node", false, "skip node related checks")
	rootCmd.Flags().BoolVar(&options.SkipDB, "skip-db", false, "skip DB related checks")
}
