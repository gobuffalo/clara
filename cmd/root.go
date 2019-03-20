package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/gobuffalo/doctor/genny/rx"
	"github.com/gobuffalo/genny"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "doctor",
	Short: "A brief description of your application",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		run := genny.WetRunner(context.Background())
		if err := run.WithNew(rx.New(&rx.Options{})); err != nil {
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
}
