package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var verbose bool
var availableSubcommands []*cobra.Command

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra-test",
	Short: "Just learning cobra",
	Long: `Cli tool with basic commands.
Built to learn cobra library.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(setVerbose, setupLogger)

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	availableSubcommands = []*cobra.Command{
		helloCmd,
		countCmd,
	}

	rootCmd.AddCommand(availableSubcommands...)
}

func setVerbose() {
	log.SetLevel(log.InfoLevel)

	if verbose {
		log.SetLevel(log.DebugLevel)
	}
}

func setupLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}
