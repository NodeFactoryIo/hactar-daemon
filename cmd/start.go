package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Register command
func init() {
	RootCmd.AddCommand(startCmd)
}

// Initialize command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts hactar deamon app",
	Long:  `....`,
	Run: func(cmd *cobra.Command, args []string) {
		onStart()
	},
}

// Start function
func onStart()  {
	fmt.Printf("Hactar started...\n")
}