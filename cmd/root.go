package cmd

import (
	"fmt"
	"os"

	"github.com/pariskwsto/cube/build"
	"github.com/spf13/cobra"
)

// versionFlag will hold the value of our persistent version flag.
var versionFlag bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cube",
	Short: "A lightweight Go CLI tool for managing VPS setup and configuration",
	Long: `Cube is a lightweight and powerful CLI tool written in Go for automating the setup and configuration of VPS servers.
It simplifies essential server management tasks, such as installing and configuring software, managing users, setting up firewalls, and optimizing security settings.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Println("cube version", build.Version)
			os.Exit(0)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
		rootCmd.Version = build.Version
	
		// Add a persistent flag for version with shorthand -v.
		rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "Print the version number and exit")
}
