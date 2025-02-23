package cmd

import (
	"fmt"
	"os"

	"github.com/pariskwsto/cube/internal/ubuntu"
	"github.com/pariskwsto/cube/internal/utils"
	"github.com/spf13/cobra"
)

// ubuntuInitialSetupCmd represents the ubuntuInitialSetup command
var ubuntuInitialSetupCmd = &cobra.Command{
	Use:   "ubuntu-initial-setup",
	Short: "Initial VPS setup as root [1] (Ubuntu 24.04.1 LTS)",
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintCommand("> Starting initial VPS setup as root...")

		err := ubuntu.InitialSetupAsRoot()
		if err != nil {
			utils.PrintError(fmt.Sprintf("Error with initial VPS setup as root: %v", err))
			os.Exit(1)
		}
		
		utils.PrintSuccess("Initial VPS setup as root completed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(ubuntuInitialSetupCmd)
}
