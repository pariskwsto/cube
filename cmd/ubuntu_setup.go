package cmd

import (
	"fmt"
	"os"

	"github.com/pariskwsto/cube/internal/linux"
	"github.com/spf13/cobra"
)

// ubuntuSetupCmd represents the ubuntuSetup command
var ubuntuSetupCmd = &cobra.Command{
	Use:   "ubuntu-setup",
	Short: "Continue VPS Server Setup with Ubuntu [2] (Ubuntu 24.04.1 LTS)",
	Run: func(cmd *cobra.Command, args []string) {
		err := linux.UpdateLinux()
		if err != nil {
			fmt.Println("Error updating system:", err)
			os.Exit(1)
		}
		fmt.Println("System update completed successfully!")
	},
}

// init initializes the cobra ubuntuSetupCmd command
func init() {
	rootCmd.AddCommand(ubuntuSetupCmd)
}
