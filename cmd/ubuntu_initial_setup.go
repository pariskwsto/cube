package cmd

import (
	"fmt"
	"os"

	"github.com/pariskwsto/cube/internal/styles"
	"github.com/pariskwsto/cube/internal/ubuntu"
	"github.com/spf13/cobra"
)

// ubuntuInitialSetupCmd represents the ubuntuInitialSetup command
var ubuntuInitialSetupCmd = &cobra.Command{
	Use:   "ubuntu-initial-setup",
	Short: "Initial VPS setup as root [1] (Ubuntu 24.04.1 LTS)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(styles.Command("> Starting initial VPS setup as root..."))

		err := ubuntu.InitialSetupAsRoot()
		if err != nil {
			fmt.Println(styles.Error(fmt.Sprintf("Error with initial VPS setup as root: %v", err)))
			os.Exit(1)
		}
		
		fmt.Println(styles.Success("Initial VPS setup as root completed successfully!"))
	},
}

func init() {
	rootCmd.AddCommand(ubuntuInitialSetupCmd)
}
