package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	// Version contains the current version.
	Version = "dev"
	// BuildDate contains a string with the build date.
	BuildDate = "unknown"
)

// Register command
func init() {
	RootCmd.AddCommand(versionCmd)
}

// Initialize command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Display version and build information about hactar deamon app.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hactar %s\n", Version)
		fmt.Printf("Built with: %s\n", runtime.Version())
	},
}
