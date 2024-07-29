package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var name string

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints greeting message",
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("hello called")
		fmt.Printf("Hello, %s!\n", name)
	},
}

func init() {
	helloCmd.Flags().StringVar(&name, "name", "World", "Name to greet")
}
