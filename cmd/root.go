package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cube",
	Short: "A lightweight Go CLI tool for managing VPS setup and configuration",
	Long: `Cube is a lightweight and powerful CLI tool written in Go for automating the setup and configuration of VPS servers.
It simplifies essential server management tasks, such as installing and configuring software, managing users, setting up firewalls, and optimizing security settings.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// init initializes the cobra root command
func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


