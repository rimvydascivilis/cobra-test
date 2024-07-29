package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var count int

var countCmd = &cobra.Command{
	Use:   "count",
	Short: "Prints numbers from 1 up to n",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("count called")
		log.Info("Counting up to ", count)

		for i := 1; i <= count; i++ {
			log.Info(i)
		}
	},
}

func init() {
	countCmd.Flags().IntVar(&count, "count", 10, "Number to count up to")
}
