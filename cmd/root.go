package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the root for all hactar commands.
var RootCmd = &cobra.Command{
	Use:           "hactar",
	Short:         "Hactar deamon app",
	SilenceErrors: true,
}