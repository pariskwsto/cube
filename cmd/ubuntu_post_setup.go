package cmd

import (
	"fmt"
	"os"

	"github.com/pariskwsto/cube/internal/linux"
	"github.com/spf13/cobra"
)

// ubuntuPostSetupCmd represents the ubuntuPostSetup command
var ubuntuPostSetupCmd = &cobra.Command{
	Use:   "ubuntu-post-setup",
	Short: "Continue VPS Setup [2] (Ubuntu 24.04.1 LTS)",
	Run: func(cmd *cobra.Command, args []string) {
		err := linux.UpdateLinux()
		if err != nil {
			fmt.Println("Error updating system:", err)
			os.Exit(1)
		}
		fmt.Println("System update completed successfully!")
	},
}

func init() {
	rootCmd.AddCommand(ubuntuPostSetupCmd)
}
