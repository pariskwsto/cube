package cmd

import (
	"fmt"
	"os"

	"github.com/pariskwsto/cube/internal/linux"
	"github.com/spf13/cobra"
)

// ubuntuSetupAsRootCmd represents the ubuntuSetupAsRoot command
var ubuntuSetupAsRootCmd = &cobra.Command{
	Use:   "ubuntu-setup-as-root",
	Short: "Initial VPS Server Setup with Ubuntu as root [1] (Ubuntu 24.04.1 LTS)",
	Run: func(cmd *cobra.Command, args []string) {
		err := linux.UpdateLinux()
		if err != nil {
			fmt.Println("Error updating system:", err)
			os.Exit(1)
		}
		fmt.Println("System update completed successfully!")
	},
}

// init initializes the cobra ubuntuSetupAsRootCmd command
func init() {
	rootCmd.AddCommand(ubuntuSetupAsRootCmd)
}
