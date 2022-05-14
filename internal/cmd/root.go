package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RootCmd create the is the root command.
var RootCmd = &cobra.Command {
	Use: "serve",
	Short: "The root command required to initialize the configuration",
}

// Execute is the entry into the cli and executing the root command.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatalf("error while executing the root command %s", err.Error())
	}
}
