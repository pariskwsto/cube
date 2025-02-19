package cmd

import (
	"fmt"
	"os"

	"github.com/pariskwsto/cube/internal/styles"
	"github.com/pariskwsto/cube/internal/ubuntu"
	"github.com/spf13/cobra"
)

// ubuntuStandardSetupCmd represents the ubuntuStandardSetup command
var ubuntuStandardSetupCmd = &cobra.Command{
	Use:   "ubuntu-standard-setup",
	Short: "Continue with standard VPS setup [2] (Ubuntu 24.04.1 LTS)",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(styles.Command("> Starting standard VPS setup..."))

		err := ubuntu.StandardSetup()
		if err != nil {
			fmt.Println(styles.Error(fmt.Sprintf("Error with standard VPS setup: %v", err)))
			os.Exit(1)
		}
		
		fmt.Println(styles.Success("Standard VPS setup completed successfully!"))
	},
}

func init() {
	rootCmd.AddCommand(ubuntuStandardSetupCmd)
}
